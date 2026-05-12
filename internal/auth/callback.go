package auth

import (
	"context"
	"crypto/subtle"
	"fmt"
	"html"
	"net"
	"net/http"
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
			fmt.Fprintf(w, "<html><body><h1>Login failed</h1><p>%s</p></body></html>", desc)
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			cs.trySendError(fmt.Errorf("no authorization code received"))
			http.Error(w, "No code received", http.StatusBadRequest)
			return
		}

		cs.trySendCode(code)
		fmt.Fprint(w, "<html><body><h1>Login successful!</h1><p>You can close this tab.</p></body></html>")
	}
}
