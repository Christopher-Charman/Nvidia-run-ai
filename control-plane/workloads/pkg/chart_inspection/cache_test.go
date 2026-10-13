/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package chart_inspection

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func cacheChart(t *testing.T) []byte {
	return cacheChartMetadata(t, "app", "1")
}

func ngcCacheSelection() Selection {
	return Selection{Repo: "https://helm.ngc.nvidia.com/nvidia", ChartName: "app", ChartVersion: "1"}
}

func TestInspectorOnlyReusesNGCRepositoryResults(t *testing.T) {
	first := cacheChart(t)
	updated := chartArchive(t,
		archiveEntry{name: "app/Chart.yaml", body: "name: app\nversion: '1'\n"},
		archiveEntry{name: "app/questions.yaml", body: "questions:\n- variable: replicas\n  type: int\n  label: Updated\n"},
	)
	for _, tc := range []struct {
		name      string
		selection Selection
		cached    bool
	}{
		{"NGC repository", ngcCacheSelection(), true},
		{"custom URL", Selection{ChartURL: "https://example.com/app.tgz"}, false},
		{"custom NGC URL", Selection{ChartURL: "https://helm.ngc.nvidia.com/app.tgz"}, false},
		{"upload", Selection{ChartID: "authorized-upload"}, false},
		{"other repository", Selection{Repo: "https://example.com/repo", ChartName: "app", ChartVersion: "1"}, false},
		{"NGC lookalike", Selection{Repo: "https://helm.ngc.nvidia.com.example.org/repo", ChartName: "app", ChartVersion: "1"}, false},
		{"HTTP NGC repository", Selection{Repo: "http://helm.ngc.nvidia.com/nvidia", ChartName: "app", ChartVersion: "1"}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) {
				calls++
				if calls == 1 {
					return first, nil
				}
				return updated, nil
			}, 0)
			before, err := inspector.Inspect(context.Background(), tc.selection, "tenant:caller", "")
			if err != nil {
				t.Fatal(err)
			}
			after, err := inspector.Inspect(context.Background(), tc.selection, "tenant:caller", "")
			if err != nil {
				t.Fatal(err)
			}
			if tc.cached {
				if calls != 1 || after.ArchiveSHA256 != before.ArchiveSHA256 {
					t.Fatal("NGC repository was not cached")
				}
			} else {
				if calls != 2 || after.ArchiveSHA256 == before.ArchiveSHA256 {
					t.Fatal("fresh source returned cached metadata")
				}
				if len(inspector.entries) != 0 {
					t.Fatal("fresh source populated cache")
				}
			}
		})
	}
}

func TestInspectorDefaultNGCCacheExpiresAfterTenMinutes(t *testing.T) {
	archive := cacheChart(t)
	calls := 0
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) { calls++; return archive, nil }, 0)
	now := time.Unix(100, 0)
	inspector.now = func() time.Time { return now }
	for _, second := range []int{0, 60, 599, 600} {
		now = time.Unix(100+int64(second), 0)
		if _, err := inspector.Inspect(context.Background(), ngcCacheSelection(), "tenant:caller", ""); err != nil {
			t.Fatal(err)
		}
		want := 1
		if second == 600 {
			want = 2
		}
		if calls != want {
			t.Fatalf("at %ds: acquisitions=%d, want=%d", second, calls, want)
		}
	}
}

func TestInspectorFreshRequestsDoNotShareAnAcquisition(t *testing.T) {
	archive := cacheChart(t)
	started := make(chan struct{}, 2)
	release := make(chan struct{})
	var releaseOnce sync.Once
	unblock := func() { releaseOnce.Do(func() { close(release) }) }
	t.Cleanup(unblock)
	done := make(chan error, 2)
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) {
		started <- struct{}{}
		<-release
		return archive, nil
	}, 0)
	for n := 0; n < 2; n++ {
		go func() {
			_, err := inspector.Inspect(context.Background(), Selection{ChartURL: "https://example.com/app.tgz"}, "tenant", "")
			done <- err
		}()
	}
	for n := 0; n < 2; n++ {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("fresh requests shared one acquisition")
		}
	}
	unblock()
	for n := 0; n < 2; n++ {
		if err := <-done; err != nil {
			t.Fatal(err)
		}
	}
}

func cacheChartMetadata(t *testing.T, name, version string) []byte {
	t.Helper()
	var output bytes.Buffer
	gz := gzip.NewWriter(&output)
	tw := tar.NewWriter(gz)
	for _, file := range []struct{ name, body string }{{"chart/Chart.yaml", "apiVersion: v2\nname: " + name + "\nversion: '" + version + "'\n"}, {"chart/questions.yaml", "questions:\n  - variable: replicas\n    label: Replicas\n    type: int\n    default: 1\n"}} {
		if err := tw.WriteHeader(&tar.Header{Name: file.name, Mode: 0600, Size: int64(len(file.body))}); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.body)); err != nil {
			t.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	return output.Bytes()
}

func TestInspectorCacheExpiryIsolationAndMutation(t *testing.T) {
	archive := cacheChart(t)
	var calls atomic.Int32
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) { calls.Add(1); return archive, nil }, time.Minute)
	now := time.Unix(100, 0)
	inspector.now = func() time.Time { return now }
	selection := ngcCacheSelection()
	first, err := inspector.Inspect(context.Background(), selection, "tenant1:caller1", "key1")
	if err != nil {
		t.Fatal(err)
	}
	if len(first.Questions) != 1 {
		t.Fatalf("fixture declaration was not parsed: %+v", first)
	}
	first.ChartName = "mutated"
	first.Questions[0]["label"] = "Changed"
	second, err := inspector.Inspect(context.Background(), selection, "tenant1:caller1", "key1")
	if err != nil {
		t.Fatal(err)
	}
	if second.ChartName != "app" || second.Questions[0]["label"] != "Replicas" || calls.Load() != 1 {
		t.Fatalf("cache mutated or missed: %+v calls=%d", second, calls.Load())
	}
	for _, scope := range []string{"tenant2:caller1", "tenant1:caller2"} {
		if _, err := inspector.Inspect(context.Background(), selection, scope, "key1"); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := inspector.Inspect(context.Background(), selection, "tenant1:caller1", "rotated-key"); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 4 {
		t.Fatalf("identity/rotation isolation failed: %d", calls.Load())
	}
	now = now.Add(time.Minute)
	if _, err := inspector.Inspect(context.Background(), selection, "tenant1:caller1", "key1"); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 5 {
		t.Fatalf("expired hit: %d", calls.Load())
	}
	inspector.mu.Lock()
	defer inspector.mu.Unlock()
	for key := range inspector.entries {
		if strings.Contains(key, "key1") || strings.Contains(key, "rotated-key") {
			t.Fatal("plaintext key cached")
		}
	}
}

func TestInspectorPreservesIntegerDefaultsAcrossCacheHits(t *testing.T) {
	archive := chartArchive(t,
		archiveEntry{name: "app/Chart.yaml", body: "name: app\nversion: '1'\n"},
		archiveEntry{name: "app/questions.yaml", body: "questions:\n  - variable: counter\n    type: int\n    label: Counter\n    default: 9007199254740993\n"},
	)
	var calls int
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) {
		calls++
		return archive, nil
	}, time.Minute)
	for n := 0; n < 2; n++ {
		result, err := inspector.Inspect(context.Background(), ngcCacheSelection(), "tenant", "")
		if err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(result.Questions[0]["default"])
		if err != nil {
			t.Fatal(err)
		}
		if string(encoded) != "9007199254740993" {
			t.Fatalf("default changed during inspection/cache: %s", encoded)
		}
	}
	if calls != 1 {
		t.Fatalf("expected the second response to use cached metadata, downloads=%d", calls)
	}
}

func TestInspectorRetriesFailures(t *testing.T) {
	archive := cacheChart(t)
	var calls int
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) {
		calls++
		if calls == 1 {
			return nil, errors.New("temporary failure")
		}
		if calls == 2 {
			return []byte("invalid archive"), nil
		}
		return archive, nil
	}, time.Minute)
	selection := ngcCacheSelection()
	for n := 0; n < 2; n++ {
		if _, err := inspector.Inspect(context.Background(), selection, "tenant", ""); err == nil {
			t.Fatal("failure accepted")
		}
	}
	if _, err := inspector.Inspect(context.Background(), selection, "tenant", ""); err != nil {
		t.Fatal(err)
	}
	if _, err := inspector.Inspect(context.Background(), selection, "tenant", ""); err != nil || calls != 3 {
		t.Fatalf("successful retry not cached: %d %v", calls, err)
	}
}

func TestInspectorSharesMissesAndKeepsWaiterAlive(t *testing.T) {
	archive := cacheChart(t)
	started := make(chan struct{})
	release := make(chan struct{})
	var calls atomic.Int32
	inspector := NewInspector(func(ctx context.Context, _ Selection, _ string) ([]byte, error) {
		if calls.Add(1) == 1 {
			close(started)
		}
		select {
		case <-release:
			return archive, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}, time.Minute)
	selection := ngcCacheSelection()
	ctx, cancel := context.WithCancel(context.Background())
	firstDone := make(chan error, 1)
	go func() { _, err := inspector.Inspect(ctx, selection, "tenant", "key"); firstDone <- err }()
	<-started
	const others = 7
	var wg sync.WaitGroup
	errs := make(chan error, others)
	for n := 0; n < others; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := inspector.Inspect(context.Background(), selection, "tenant", "key")
			if err == nil && r.ChartName != "app" {
				err = errors.New("wrong result")
			}
			errs <- err
		}()
	}
	deadline := time.Now().Add(time.Second)
	for {
		inspector.mu.Lock()
		waiters := 0
		for _, f := range inspector.flights {
			waiters = f.waiters
		}
		inspector.mu.Unlock()
		if waiters == 8 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("waiters did not attach")
		}
		time.Sleep(time.Millisecond)
	}
	cancel()
	if err := <-firstDone; !errors.Is(err, context.Canceled) {
		t.Fatalf("leader cancellation lost %v", err)
	}
	close(release)
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	if calls.Load() != 1 {
		t.Fatalf("shared %d downloads", calls.Load())
	}
}

func TestInspectorCancelsAbandonedFlightAndRetries(t *testing.T) {
	archive := cacheChart(t)
	started := make(chan struct{})
	abandoned := make(chan struct{})
	var calls atomic.Int32
	inspector := NewInspector(func(ctx context.Context, _ Selection, _ string) ([]byte, error) {
		if calls.Add(1) == 1 {
			close(started)
			<-ctx.Done()
			close(abandoned)
			return nil, ctx.Err()
		}
		return archive, nil
	}, time.Minute)
	selection := ngcCacheSelection()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { _, err := inspector.Inspect(ctx, selection, "tenant", ""); done <- err }()
	<-started
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	select {
	case <-abandoned:
	case <-time.After(time.Second):
		t.Fatal("abandoned download survived")
	}
	if _, err := inspector.Inspect(context.Background(), selection, "tenant", ""); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 2 {
		t.Fatalf("cancelled miss reused %d", calls.Load())
	}
}

func TestInspectorDoesNotShareDifferentSelections(t *testing.T) {
	archive := cacheChart(t)
	var calls atomic.Int32
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) { calls.Add(1); return archive, nil }, time.Minute)
	for _, sel := range []Selection{{ChartID: "one"}, {ChartID: "two"}} {
		if _, err := inspector.Inspect(context.Background(), sel, "tenant", ""); err != nil {
			t.Fatal(err)
		}
	}
	if calls.Load() != 2 {
		t.Fatal("distinct selections shared")
	}
	if _, err := inspector.Inspect(context.Background(), Selection{Repo: "https://example.com", ChartName: "wrong", ChartVersion: "1"}, "tenant", ""); err == nil {
		t.Fatal("wrong package identity accepted")
	}
}

func TestInspectorBoundsActiveAndCachedWork(t *testing.T) {
	archive := cacheChart(t)
	release := make(chan struct{})
	started := make(chan struct{}, 4)
	inspector := NewInspector(func(ctx context.Context, _ Selection, _ string) ([]byte, error) {
		started <- struct{}{}
		select {
		case <-release:
			return archive, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}, time.Minute)
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			_, _ = inspector.Inspect(context.Background(), Selection{ChartID: string(rune('a' + n))}, "tenant", "")
		}(n)
	}
	for n := 0; n < 4; n++ {
		<-started
	}
	if _, err := inspector.Inspect(context.Background(), Selection{ChartID: "fifth"}, "tenant", ""); err == nil {
		t.Fatal("unbounded active work")
	}
	close(release)
	wg.Wait()
	inspector.maxEntries = 2
	inspector.fetch = func(context.Context, Selection, string) ([]byte, error) { return archive, nil }
	for _, id := range []string{"x", "y", "z"} {
		selection := ngcCacheSelection()
		selection.Repo += "/" + id
		if _, err := inspector.Inspect(context.Background(), selection, "tenant", ""); err != nil {
			t.Fatal(err)
		}
	}
	inspector.mu.Lock()
	count := len(inspector.entries)
	inspector.mu.Unlock()
	if count != 2 {
		t.Fatalf("unbounded entries %d", count)
	}
	inspector.maxRecordBytes = 1
	if _, err := inspector.Inspect(context.Background(), ngcCacheSelection(), "tenant", ""); err != nil {
		t.Fatal(err)
	}
	inspector.mu.Lock()
	defer inspector.mu.Unlock()
	if inspector.cachedBytes > inspector.maxCacheBytes {
		t.Fatal("unbounded cached bytes")
	}
}

func TestInspectorRejectsIdentityMismatchWithoutCaching(t *testing.T) {
	wrong, correct := cacheChart(t), cacheChartMetadata(t, "app", "2")
	var calls int
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) {
		calls++
		if calls < 3 {
			return wrong, nil
		}
		return correct, nil
	}, time.Minute)
	selection := Selection{Repo: "https://helm.ngc.nvidia.com/nvidia", ChartName: "app", ChartVersion: "2"}
	for n := 0; n < 2; n++ {
		_, err := inspector.Inspect(context.Background(), selection, "tenant", "")
		var typed *Error
		if !errors.As(err, &typed) || typed.StatusCode != 422 || typed.Code != "chart_identity_mismatch" {
			t.Fatalf("wrong mismatch error: %v", err)
		}
	}
	result, err := inspector.Inspect(context.Background(), selection, "tenant", "")
	if err != nil || result.ChartVersion != "2" || calls != 3 {
		t.Fatalf("failed identity was cached: %+v %v calls=%d", result, err, calls)
	}
	if _, err := inspector.Inspect(context.Background(), selection, "tenant", ""); err != nil || calls != 3 {
		t.Fatalf("correct identity not cached: %v calls=%d", err, calls)
	}
}

func TestInspectorServesUploadedCallbackFresh(t *testing.T) {
	archive := cacheChart(t)
	var calls int
	inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) {
		t.Error("remote fetch used for authorized upload")
		return nil, errors.New("wrong source")
	}, time.Minute)
	uploaded := func(context.Context, Selection, string) ([]byte, error) { calls++; return archive, nil }
	for n := 0; n < 2; n++ {
		result, err := inspector.InspectWithFetch(context.Background(), Selection{ChartID: "authorized-upload"}, "tenant", "", uploaded)
		if err != nil || result.ChartName != "app" {
			t.Fatalf("uncacheable result lost: %+v %v", result, err)
		}
	}
	if calls != 2 {
		t.Fatalf("upload was cached: %d", calls)
	}
	if len(inspector.entries) != 0 {
		t.Fatal("upload populated cache")
	}
}

func TestInspectorServesNGCResultExceedingCacheBudgets(t *testing.T) {
	archive := cacheChart(t)
	for _, budget := range []string{"record", "total"} {
		t.Run(budget, func(t *testing.T) {
			calls := 0
			inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) {
				calls++
				return archive, nil
			}, 0)
			if budget == "record" {
				inspector.maxRecordBytes = 1
			} else {
				inspector.maxCacheBytes = 1
			}
			for n := 0; n < 2; n++ {
				result, err := inspector.Inspect(context.Background(), ngcCacheSelection(), "tenant", "")
				if err != nil || result.ChartName != "app" {
					t.Fatalf("uncacheable result lost: %+v %v", result, err)
				}
			}
			if calls != 2 || len(inspector.entries) != 0 {
				t.Fatalf("cache %s budget ignored: acquisitions=%d entries=%d", budget, calls, len(inspector.entries))
			}
		})
	}
}
