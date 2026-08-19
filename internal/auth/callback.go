package auth

import (
	"context"
	"crypto/subtle"
	"fmt"
	"html"
	"net"
	"net/http"
	"strings"
	"time"
)

type callbackServer struct {
	listener net.Listener
	server   *http.Server
	codeCh   chan string
	errCh    chan error
	done     chan struct{}
}

func newCallbackServer(expectedState string) (*callbackServer, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	cs := &callbackServer{
		listener: listener,
		codeCh:   make(chan string, 1),
		errCh:    make(chan error, 1),
		done:     make(chan struct{}),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", cs.handleCallback(expectedState))

	cs.server = &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	return cs, nil
}

func (cs *callbackServer) RedirectURI() string {
	port := cs.listener.Addr().(*net.TCPAddr).Port
	return fmt.Sprintf("http://127.0.0.1:%d/callback", port)
}

func (cs *callbackServer) Start() {
	go cs.server.Serve(cs.listener)
}

func (cs *callbackServer) Close() {
	cs.server.Close()
}

func (cs *callbackServer) WaitForCode(ctx context.Context, timeout time.Duration) (string, error) {
	defer close(cs.done)
	select {
	case code := <-cs.codeCh:
		return code, nil
	case err := <-cs.errCh:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(timeout):
		return "", fmt.Errorf("authentication timed out")
	}
}

func (cs *callbackServer) trySendCode(code string) {
	select {
	case cs.codeCh <- code:
	case <-cs.done:
	}
}

func (cs *callbackServer) trySendError(err error) {
	select {
	case cs.errCh <- err:
	case <-cs.done:
	}
}

const callbackPage = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="color-scheme" content="light dark">
<title>Nimble CLI</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet">
<style>
  * { margin: 0; padding: 0; box-sizing: border-box; }
  body { font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif; min-height: 100vh; display: flex; align-items: center; justify-content: center; background: #08090a; color: #f7f8f8; }
  .card { background: rgb(25, 26, 27); border: 1px solid rgba(255, 255, 255, 0.08); border-radius: 16px; backdrop-filter: blur(24px) saturate(160%%); padding: 48px 40px; max-width: 520px; width: 90%%; text-align: center; display: flex; flex-direction: column; align-items: center; gap: 20px; }
  .logo { width: 36px; height: 30px; color: #f7f8f8; }
  .divider { width: 40px; height: 1px; background: rgba(255, 255, 255, 0.08); }
  .status-icon { width: 56px; height: 56px; border-radius: 50%%; display: flex; align-items: center; justify-content: center; }
  .status-icon.success { background: rgba(67, 160, 71, 0.15); color: #43a047; }
  .status-icon.error { background: rgba(211, 47, 47, 0.15); color: #d32f2f; }
  .status-icon svg { width: 28px; height: 28px; }
  h1 { font-size: 20px; font-weight: 600; letter-spacing: -0.01em; }
  .message { font-size: 14px; color: rgba(247, 248, 248, 0.6); line-height: 1.6; max-width: 320px; }
  .app-name { font-size: 12px; font-weight: 500; color: rgba(247, 248, 248, 0.4); letter-spacing: 0.05em; text-transform: uppercase; }
</style>
</head>
<body>
<div class="card">
  <svg viewBox="0 0 94 79" class="logo"><path d="M93.3311 31.4141V78.834L39.916 37.8975V78.8408L0 47.3652V0L53.4277 41.1123V0L93.3311 31.4141Z" fill="currentColor"/></svg>
  <div class="divider"></div>
  {{CONTENT}}
  <p class="app-name">Nimble CLI</p>
</div>
</body>
</html>`

var successPage = strings.Replace(callbackPage, "{{CONTENT}}", `
  <div class="status-icon success"><svg viewBox="0 -960 960 960" fill="currentColor"><path d="M382-240 154-468l57-57 171 171 367-367 57 57-424 424Z"/></svg></div>
  <h1>Login successful</h1>
  <p class="message">You can close this tab and return to the terminal.</p>
`, 1)

const errorPageTemplate = `
  <div class="status-icon error"><svg viewBox="0 -960 960 960" fill="currentColor"><path d="m256-200-56-56 224-224-224-224 56-56 224 224 224-224 56 56-224 224 224 224-56 56-224-224-224 224Z"/></svg></div>
  <h1>Login failed</h1>
  <p class="message">%s</p>
`

func (cs *callbackServer) handleCallback(expectedState string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		if subtle.ConstantTimeCompare([]byte(state), []byte(expectedState)) != 1 {
			cs.trySendError(fmt.Errorf("state mismatch"))
			http.Error(w, "State mismatch", http.StatusBadRequest)
			return
		}

		if errParam := r.URL.Query().Get("error"); errParam != "" {
			desc := html.EscapeString(r.URL.Query().Get("error_description"))
			cs.trySendError(fmt.Errorf("authorization denied: %s (%s)", errParam, desc))
			errorContent := fmt.Sprintf(errorPageTemplate, desc)
			fmt.Fprint(w, strings.Replace(callbackPage, "{{CONTENT}}", errorContent, 1))
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			cs.trySendError(fmt.Errorf("no authorization code received"))
			http.Error(w, "No code received", http.StatusBadRequest)
			return
		}

		cs.trySendCode(code)
		fmt.Fprint(w, successPage)
	}
}
