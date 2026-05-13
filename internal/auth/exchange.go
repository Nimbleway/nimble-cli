package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type oauthEndpoints struct {
	authorizationEndpoint string
	tokenEndpoint         string
	registrationEndpoint  string
}

func discoverEndpoints(ctx context.Context, baseURL string) (*oauthEndpoints, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/.well-known/oauth-authorization-server", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		AuthorizationEndpoint string `json:"authorization_endpoint"`
		TokenEndpoint         string `json:"token_endpoint"`
		RegistrationEndpoint  string `json:"registration_endpoint"`
	}
	if err := doJSON(req, &raw); err != nil {
		return nil, err
	}

	return &oauthEndpoints{
		authorizationEndpoint: raw.AuthorizationEndpoint,
		tokenEndpoint:         raw.TokenEndpoint,
		registrationEndpoint:  raw.RegistrationEndpoint,
	}, nil
}

func registerClient(ctx context.Context, endpoint, redirectURI string) (string, error) {
	body := map[string]interface{}{
		"client_name":                "Nimble CLI",
		"redirect_uris":             []string{redirectURI},
		"grant_types":               []string{"authorization_code"},
		"response_types":            []string{"code"},
		"token_endpoint_auth_method": "none",
	}
	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	var resp struct {
		ClientID string `json:"client_id"`
	}
	if err := doJSON(req, &resp); err != nil {
		return "", err
	}
	return resp.ClientID, nil
}

func buildAuthorizeURL(endpoint, clientID, redirectURI, challenge, state string) string {
	v := url.Values{
		"client_id":             {clientID},
		"redirect_uri":          {redirectURI},
		"response_type":         {"code"},
		"code_challenge":        {challenge},
		"code_challenge_method": {"S256"},
		"state":                 {state},
		"scope":                 {"openid offline_access nimble:read nimble:write"},
	}
	return endpoint + "?" + v.Encode()
}

func exchangeCode(ctx context.Context, endpoint, clientID, code, verifier, redirectURI string) (string, error) {
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"code_verifier": {verifier},
		"client_id":     {clientID},
		"redirect_uri":  {redirectURI},
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var resp struct {
		AccessToken string `json:"access_token"`
	}
	if err := doJSON(req, &resp); err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}
