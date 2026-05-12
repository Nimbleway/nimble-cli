package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func decodeBody(resp *http.Response, dest interface{}) error {
	return json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(dest)
}
