package identity_manager_client

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// defaultRefreshBuffer is how long before expiry we consider the token expired and refresh.
const defaultRefreshBuffer = 30 * time.Second

// AuthnManager manages tokens and implements http.RoundTripper to inject auth on each request.
// It does not hold the returned client to avoid a cycle (client's transport is the manager).
type AuthnManager struct {
	idmClient     IIdm
	clientId      string
	clientSecret  string
	accessToken   string
	tokenExpiresAt time.Time
	mu            sync.Mutex
	refreshBuffer time.Duration
}

type AuthnManagerConfig struct {
	ClientId      string
	ClientSecret  string
	IdmClient     IIdm
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

	_, err := authnManager.setToken(ctx)
	if err != nil {
		return nil, err
	}
	return authnManager, nil
}

func (a *AuthnManager) setToken(ctx context.Context) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	tokenResponse, err := a.idmClient.ClientCredentialsLogin(ctx, a.clientId, a.clientSecret)
	if err != nil {
		return "", err
	}
	a.accessToken = tokenResponse.GetAccessToken()
	a.tokenExpiresAt = time.Now().Add(time.Duration(tokenResponse.GetExpiresIn()) * time.Second)
	return a.accessToken, nil
}

func (a *AuthnManager) Token(ctx context.Context) (string, error) {
	if a.accessToken == "" || time.Now().After(a.tokenExpiresAt.Add(-defaultRefreshBuffer)) {
		return a.setToken(ctx)
	}
	return a.accessToken, nil
}


// RoundTrip sets the Authorization header, and delegates to http.DefaultTransport.
func (a *AuthnManager) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := a.Token(req.Context())
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	return http.DefaultTransport.RoundTrip(req)
}

// HttpClient returns an HTTP client that uses this manager as its transport.
// The manager does not retain a reference to the returned client (avoids cycle).
func (a *AuthnManager) HttpClient(ctx context.Context) (*http.Client, error) {
	return &http.Client{Transport: a}, nil
}