package auth

import (
	"encoding/base64"
	"testing"
)

func jwtWithPayload(payload string) string {
	enc := func(s string) string { return base64.RawURLEncoding.EncodeToString([]byte(s)) }
	return enc(`{"alg":"HS256","typ":"JWT"}`) + "." + enc(payload) + "." + enc("signature")
}

func TestTokenUsername(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  string
	}{
		{"username claim", jwtWithPayload(`{"username":"omer@example.com"}`), "omer@example.com"},
		{"claim absent", jwtWithPayload(`{"account_name":"acct"}`), ""},
		{"payload not json", jwtWithPayload(`not json`), ""},
		{"not a jwt", "opaque-token", ""},
		{"too few segments", "header.payload", ""},
		{"payload not base64url", "header.!!!.signature", ""},
		{"empty", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tokenUsername(tc.token); got != tc.want {
				t.Errorf("tokenUsername() = %q, want %q", got, tc.want)
			}
		})
	}
}
