package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func decodeBody(resp *http.Response, dest interface{}) error {
	return json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(dest)
}

func doJSON(req *http.Request, dest interface{}) error {
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s %s returned status %d", req.Method, req.URL.Path, resp.StatusCode)
	}

	return decodeBody(resp, dest)
}
