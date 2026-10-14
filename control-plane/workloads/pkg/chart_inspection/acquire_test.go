/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package chart_inspection

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

type inspectionResolver func(context.Context, string) ([]net.IPAddr, error)

func (r inspectionResolver) LookupIPAddr(ctx context.Context, host string) ([]net.IPAddr, error) {
	return r(ctx, host)
}

func testRemoteFetcher(t *testing.T, handler http.HandlerFunc) (*remoteFetcher, string, *[]string) {
	t.Helper()
	server := httptest.NewTLSServer(handler)
	t.Cleanup(server.Close)
	address := server.Listener.Addr().String()
	_, port, _ := net.SplitHostPort(address)
	var mu sync.Mutex
	targets := []string{}
	resolver := inspectionResolver(func(context.Context, string) ([]net.IPAddr, error) {
		return []net.IPAddr{{IP: net.ParseIP("8.8.8.8")}}, nil
	})
	dial := func(ctx context.Context, network, target string) (net.Conn, error) {
		mu.Lock()
		targets = append(targets, target)
		mu.Unlock()
		return (&net.Dialer{}).DialContext(ctx, network, address)
	}
	fetcher := newRemoteFetcher(resolver, dial)
	roots := x509.NewCertPool()
	roots.AddCert(server.Certificate())
	fetcher.tlsConfig = &tls.Config{RootCAs: roots, ServerName: server.Certificate().DNSNames[0], MinVersion: tls.VersionTLS12}
	return fetcher, "https://helm.ngc.nvidia.com:" + port, &targets
}

func TestFetchRemoteDirectAndRepository(t *testing.T) {
	var paths []string
	fetcher, base, targets := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		username, password, ok := r.BasicAuth()
		if !ok || username != "$oauthtoken" || password != "synthetic-test-key" {
			t.Errorf("missing scoped NGC authentication")
		}
		switch r.URL.Path {
		case "/repo/index.yaml":
			io.WriteString(w, "apiVersion: v1\nentries:\n  app:\n    - name: app\n      version: '1.0'\n      urls: [wrong.tgz]\n    - name: app\n      version: '2.0'\n      urls: ['../app.tgz?signature=synthetic']\n")
		case "/app.tgz":
			io.WriteString(w, "chart-archive")
		default:
			http.NotFound(w, r)
		}
	})
	got, err := fetcher.fetch(context.Background(), Selection{Repo: base + "/repo", ChartName: "app", ChartVersion: "2.0"}, "synthetic-test-key")
	if err != nil || string(got) != "chart-archive" {
		t.Fatalf("got %q, %v", got, err)
	}
	if strings.Join(paths, ",") != "/repo/index.yaml,/app.tgz" {
		t.Fatalf("paths %v", paths)
	}
	for _, address := range *targets {
		if !strings.HasPrefix(address, "8.8.8.8:") {
			t.Fatalf("validated IP not pinned: %s", address)
		}
	}
	got, err = fetcher.fetch(context.Background(), Selection{ChartURL: base + "/app.tgz"}, "synthetic-test-key")
	if err != nil || string(got) != "chart-archive" {
		t.Fatalf("direct got %q, %v", got, err)
	}
}

func TestFetchRemoteBoundsExpandedIndexAliases(t *testing.T) {
	index := "entries:\n  app: [{name: app, version: '1', urls: [app.tgz]}]\nlarge: &large " + strings.Repeat("x", 128<<10) + "\nrepeated: [*large" + strings.Repeat(", *large", 99) + "]\n"
	fetcher, base, _ := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/repo/index.yaml" {
			io.WriteString(w, index)
			return
		}
		io.WriteString(w, "chart-archive")
	})
	_, err := fetcher.fetch(context.Background(), Selection{Repo: base + "/repo", ChartName: "app", ChartVersion: "1"}, "")
	var typed *Error
	if !errors.As(err, &typed) || typed.Code != "invalid_chart_index" {
		t.Fatalf("expanded repository aliases were accepted: %v", err)
	}
}

func TestFetchRemoteRejectsUnsafeSources(t *testing.T) {
	for _, raw := range []string{"file:///etc/passwd", "oci://registry.example.com/app", "http://127.0.0.1/a.tgz", "https://[::1]/a.tgz", "http://169.254.169.254/a", "https://user:password@example.com/a", "https://example.com/a#fragment"} {
		t.Run(raw, func(t *testing.T) {
			_, err := FetchRemote(context.Background(), Selection{ChartURL: raw}, "")
			if err == nil {
				t.Fatal("unsafe source accepted")
			}
			var typed *Error
			if !errors.As(err, &typed) || typed.StatusCode != http.StatusBadRequest {
				t.Fatalf("untyped error: %v", err)
			}
			if strings.Contains(err.Error(), "password") {
				t.Fatal("embedded credentials leaked")
			}
		})
	}
}

func TestFetchRemoteBlocksPrivateDNSBeforeDial(t *testing.T) {
	dialed := false
	fetcher := newRemoteFetcher(inspectionResolver(func(context.Context, string) ([]net.IPAddr, error) {
		return []net.IPAddr{{IP: net.ParseIP("8.8.8.8")}, {IP: net.ParseIP("10.1.2.3")}}, nil
	}), func(context.Context, string, string) (net.Conn, error) {
		dialed = true
		return nil, errors.New("unexpected dial")
	})
	_, err := fetcher.fetch(context.Background(), Selection{ChartURL: "https://registry.example.com/app.tgz"}, "")
	if err == nil || dialed {
		t.Fatalf("private DNS accepted or dialed: %v, %v", err, dialed)
	}
}

func TestFetchRemoteRedirectRechecksDNSAndStripsCredential(t *testing.T) {
	fetcher, base, _ := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/first" {
			http.Redirect(w, r, "https://cdn.example.com:"+strings.Split(r.Host, ":")[1]+"/archive?signature=synthetic", http.StatusFound)
			return
		}
		if r.Header.Get("Authorization") != "" {
			t.Error("NGC credential leaked to CDN")
		}
		io.WriteString(w, "archive")
	})
	got, err := fetcher.fetch(context.Background(), Selection{ChartURL: base + "/first"}, "synthetic-test-key")
	if err != nil || string(got) != "archive" {
		t.Fatalf("signed CDN redirect: %q %v", got, err)
	}
	fetcher.resolver = inspectionResolver(func(_ context.Context, host string) ([]net.IPAddr, error) {
		if host == "cdn.example.com" {
			return []net.IPAddr{{IP: net.ParseIP("192.168.1.1")}}, nil
		}
		return []net.IPAddr{{IP: net.ParseIP("8.8.8.8")}}, nil
	})
	_, err = fetcher.fetch(context.Background(), Selection{ChartURL: base + "/first"}, "synthetic-test-key")
	if err == nil {
		t.Fatal("redirect to private DNS accepted")
	}
	if strings.Contains(err.Error(), "signature") || strings.Contains(err.Error(), "synthetic-test-key") {
		t.Fatal("secret URL or credential leaked")
	}
}

func TestFetchRemoteDropsCredentialForCrossOriginIndexLink(t *testing.T) {
	fetcher, base, _ := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/index.yaml" {
			io.WriteString(w, "entries:\n  app:\n    - version: '1'\n      urls: ['https://cdn.example.com:"+strings.Split(r.Host, ":")[1]+"/chart.tgz']\n")
			return
		}
		if r.Header.Get("Authorization") != "" {
			t.Error("credential leaked through package URL")
		}
		io.WriteString(w, "archive")
	})
	got, err := fetcher.fetch(context.Background(), Selection{Repo: base, ChartName: "app", ChartVersion: "1"}, "synthetic-test-key")
	if err != nil || string(got) != "archive" {
		t.Fatalf("got %q %v", got, err)
	}
}

func TestFetchRemoteValidatesSelectionIndexAndLimits(t *testing.T) {
	cases := []struct {
		name, body string
		selection  Selection
		wantStatus int
	}{
		{name: "duplicate index keys", body: "entries: {}\nentries: {}\n", wantStatus: 422},
		{name: "missing exact version", body: "entries:\n  app:\n    - version: '2'\n      urls: [chart.tgz]\n", wantStatus: 422},
		{name: "index too large", body: strings.Repeat("x", (1<<20)+1), wantStatus: 422},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fetcher, base, _ := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, tc.body) })
			_, err := fetcher.fetch(context.Background(), Selection{Repo: base, ChartName: "app", ChartVersion: "1"}, "")
			var typed *Error
			if !errors.As(err, &typed) || typed.StatusCode != tc.wantStatus {
				t.Fatalf("wrong failure %v", err)
			}
		})
	}
	for _, selection := range []Selection{{}, {Repo: "https://example.com", ChartName: "app"}, {ChartURL: "https://example.com/a", Repo: "https://example.com", ChartName: "app", ChartVersion: "1"}, {ChartID: "stored"}} {
		if _, err := FetchRemote(context.Background(), selection, ""); err == nil {
			t.Fatalf("invalid selection accepted %+v", selection)
		}
	}
	if _, err := FetchRemote(context.Background(), Selection{ChartURL: "https://notngc.nvidia.com/a"}, "synthetic-test-key"); err == nil {
		t.Fatal("key accepted for non-NGC host")
	}
}

func TestFetchRemoteBoundsArchiveAndRedactsFailures(t *testing.T) {
	fetcher, base, _ := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/large" {
			w.Header().Set("Content-Length", "33554433")
			return
		}
		w.WriteHeader(http.StatusForbidden)
	})
	for _, path := range []string{"/large", "/forbidden?signature=sensitive"} {
		_, err := fetcher.fetch(context.Background(), Selection{ChartURL: base + path}, "")
		if err == nil || strings.Contains(err.Error(), "sensitive") || strings.Contains(err.Error(), base) {
			t.Fatalf("missing bound or unredacted error: %v", err)
		}
	}
}

func TestFetchRemoteCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := FetchRemote(ctx, Selection{ChartURL: "https://example.com/a"}, "")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation lost: %v", err)
	}
}

func TestFetchRemotePreservesEscapedRepositoryPath(t *testing.T) {
	fetcher, base, _ := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/repo/team/index.yaml" {
			if r.URL.EscapedPath() != "/repo%2Fteam/index.yaml" {
				t.Errorf("repository path changed: %s", r.URL.EscapedPath())
			}
			io.WriteString(w, "entries:\n  app:\n    - version: '1'\n      urls: [/archive.tgz]\n")
			return
		}
		io.WriteString(w, "archive")
	})
	if _, err := fetcher.fetch(context.Background(), Selection{Repo: base + "/repo%2Fteam", ChartName: "app", ChartVersion: "1"}, ""); err != nil {
		t.Fatal(err)
	}
}

func TestFetchRemoteUpstreamHTTPFailuresReturnBadGateway(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError, http.StatusServiceUnavailable} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			fetcher, base, _ := testRemoteFetcher(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = io.WriteString(w, "upstream-sensitive-body")
			})
			for _, selection := range []Selection{
				{ChartURL: base + "/chart.tgz?signature=upstream-sensitive-query"},
				{Repo: base, ChartName: "app", ChartVersion: "1"},
			} {
				_, err := fetcher.fetch(context.Background(), selection, "synthetic-test-key")
				var sourceErr *Error
				if !errors.As(err, &sourceErr) || sourceErr.StatusCode != http.StatusBadGateway || sourceErr.Code != "chart_download_failed" {
					t.Fatalf("upstream HTTP %d returned %v", status, err)
				}
				if strings.Contains(err.Error(), "upstream-sensitive") || strings.Contains(err.Error(), "synthetic-test-key") {
					t.Fatalf("upstream failure exposed source material: %v", err)
				}
			}
		})
	}
}
