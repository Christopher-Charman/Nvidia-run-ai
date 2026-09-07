package identity_manager_client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	defaultRefreshBuffer = 30 * time.Second
	authHeaderKey        = "Authorization"
	bearerPrefix         = "Bearer "
)

type AuthnManager struct {
	idmClient      IIdmExternalClient
	clientId       string
	clientSecret   string
	accessToken    string
	tokenExpiresAt time.Time
	mu             sync.RWMutex
	refreshBuffer  time.Duration
}

type AuthnManagerConfig struct {
	ClientId      string
	ClientSecret  string
	IdmClient     IIdmExternalClient
	RefreshBuffer time.Duration
}

func NewAuthnManager(ctx context.Context, config AuthnManagerConfig) (*AuthnManager, error) {
	if config.RefreshBuffer == 0 {
		config.RefreshBuffer = defaultRefreshBuffer
	}
	authnManager := &AuthnManager{
		clientId:      config.ClientId,
		clientSecret:  config.ClientSecret,
		idmClient:     config.IdmClient,
		refreshBuffer: config.RefreshBuffer,
	}

	authnManager.mu.Lock()
	_, err := authnManager.setToken(ctx)
	authnManager.mu.Unlock()
	if err != nil {
		return nil, err
	}
	return authnManager, nil
}


func (a *AuthnManager) setToken(ctx context.Context) (string, error) {
	tokenResponse, err := a.idmClient.ClientCredentialsLogin(ctx, a.clientId, a.clientSecret)
	if err != nil {
		return "", err
	}
	a.accessToken = tokenResponse.GetAccessToken()
	a.tokenExpiresAt = time.Now().Add(time.Duration(tokenResponse.GetExpiresIn()) * time.Second)
	return a.accessToken, nil
}

func (a *AuthnManager) Token(ctx context.Context) (string, error) {
	a.mu.RLock()
	token := a.accessToken
	expiresAt := a.tokenExpiresAt
	buffer := a.refreshBuffer
	a.mu.RUnlock()

	if token != "" && !time.Now().After(expiresAt.Add(-buffer)) {
		return token, nil
	}

	
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.accessToken != "" && !time.Now().After(a.tokenExpiresAt.Add(-a.refreshBuffer)) {
		return a.accessToken, nil
	}
	return a.setToken(ctx)
}

func (a *AuthnManager) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := a.Token(req.Context())
	if err != nil {
		return nil, err
	}

	clonedReq := req.Clone(req.Context())
	clonedReq.Header.Set(authHeaderKey, bearerPrefix+token)
	resp, err := http.DefaultTransport.RoundTrip(clonedReq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return a.retryWithFreshToken(req, resp)
	}
	return resp, nil
}

func (a *AuthnManager) retryWithFreshToken(req *http.Request, staleResp *http.Response) (*http.Response, error) {
	_, _ = io.Copy(io.Discard, staleResp.Body)
	staleResp.Body.Close()

	a.mu.Lock()
	token, err := a.setToken(req.Context())
	a.mu.Unlock()
	if err != nil {
		return nil, err
	}
	retryReq := req.Clone(req.Context())
	if req.Body != nil {
		if req.GetBody == nil {
			return nil, fmt.Errorf("authn_manager: cannot retry request with non-replayable body")
		}
		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		retryReq.Body = body
	}
	retryReq.Header.Set(authHeaderKey, bearerPrefix+token)
	return http.DefaultTransport.RoundTrip(retryReq)
}

func (a *AuthnManager) HttpClient() *http.Client {
	return &http.Client{Transport: a}
}
