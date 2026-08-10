package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestCLIKeyNameCollidesOnTruncation records a limit of name-based scoping: two
// usernames that differ only past the truncation point produce the same key
// name. The name cannot distinguish them, which is why cleanup also consults
// created_by.
func TestCLIKeyNameCollidesOnTruncation(t *testing.T) {
	const host = "sapman-ThinkPad-X1-Carbon-Gen-13"
	long := strings.Repeat("a", 40)

	first := cliKeyNameFor(long+"lice@nimbleway.com", host)
	second := cliKeyNameFor(long+"ob@nimbleway.com", host)
	if first != second {
		t.Skipf("names no longer collide (%q vs %q); created_by scoping is then belt and braces", first, second)
	}
}

// keyListServer serves a fixed key list and records every path it is asked to
// delete.
func keyListServer(t *testing.T, keys []apiKeyEntry) (*httptest.Server, *[]string) {
	t.Helper()
	var deleted []string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api/v1/account/api-key" && r.Method == http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(keys); err != nil {
				t.Errorf("encoding key list: %v", err)
			}
		case strings.HasPrefix(r.URL.Path, "/api/v1/account/api-key/") && r.Method == http.MethodDelete:
			deleted = append(deleted, strings.TrimPrefix(r.URL.Path, "/api/v1/account/api-key/"))
			w.WriteHeader(http.StatusOK)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	return srv, &deleted
}

// TestListCLIKeysExcludesOtherUsers covers the case name scoping cannot: two
// users sharing a key name because their usernames were truncated to the same
// prefix. Only the caller's own key may be returned for deletion.
func TestListCLIKeysExcludesOtherUsers(t *testing.T) {
	const name = "CLI (shared @ host)"
	srv, _ := keyListServer(t, []apiKeyEntry{
		{GUID: "mine", KeyName: name, CreatedBy: "alice@nimbleway.com"},
		{GUID: "theirs", KeyName: name, CreatedBy: "bob@nimbleway.com"},
		{GUID: "other-name", KeyName: "CLI (alice @ elsewhere)", CreatedBy: "alice@nimbleway.com"},
	})

	guids, err := listCLIKeys(context.Background(), srv.URL, "token", name, "alice@nimbleway.com")
	if err != nil {
		t.Fatalf("listCLIKeys: %v", err)
	}

	if len(guids) != 1 || guids[0] != "mine" {
		t.Errorf("listCLIKeys = %v, want only [mine]", guids)
	}
}

// TestListCLIKeysWithoutCreatedBy keeps cleanup working against a server that
// does not report created_by, where the name is all there is to go on.
func TestListCLIKeysWithoutCreatedBy(t *testing.T) {
	const name = "CLI (alice @ host)"
	srv, _ := keyListServer(t, []apiKeyEntry{
		{GUID: "mine", KeyName: name},
		{GUID: "other-name", KeyName: "CLI (bob @ host)"},
	})

	guids, err := listCLIKeys(context.Background(), srv.URL, "token", name, "alice@nimbleway.com")
	if err != nil {
		t.Fatalf("listCLIKeys: %v", err)
	}

	if len(guids) != 1 || guids[0] != "mine" {
		t.Errorf("listCLIKeys = %v, want only [mine]", guids)
	}
}

// TestIsKeyLimitResponse covers the envelopes and phrasings the key limit
// arrives in. Matching none of them disables limit recovery without any visible
// symptom beyond a bare 403.
func TestIsKeyLimitResponse(t *testing.T) {
	tests := []struct {
		name string
		body string
		want bool
	}{
		{"proxit error envelope", `{"error":"max api keys limit reached"}`, true},
		{"message envelope", `{"message":"API key limit reached"}`, true},
		{"fastapi detail envelope", `{"detail":"API key limit reached"}`, true},
		{"mixed case", `{"error":"MAX API KEYS LIMIT REACHED"}`, true},
		{"unrelated forbidden", `{"error":"forbidden"}`, false},
		{"read only user", `{"message":"user is read only"}`, false},
		{"empty body", ``, false},
		{"not json", `<html>403</html>`, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader(tc.body)),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
			}

			if got := isKeyLimitResponse(resp); got != tc.want {
				t.Errorf("isKeyLimitResponse(%s) = %v, want %v", tc.body, got, tc.want)
			}
		})
	}
}
