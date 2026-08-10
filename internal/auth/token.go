package auth

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// tokenUsername reads the "username" claim from an access token, which is the
// value proxit records as an API key's created_by. It is used only to narrow
// which keys a login may revoke, so the signature is not checked: the server
// remains the authority on what this token is allowed to do, and a claim that
// cannot be read simply means no key is treated as ours.
//
// An empty string is returned for any token that is not a readable JWT.
func tokenUsername(token string) string {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return ""
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}

	var claims struct {
		Username string `json:"username"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return ""
	}
	return claims.Username
}
