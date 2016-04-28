package testutil

import (
	"net/http"
	"net/http/httptest"
)

func MockHTTPServer(ct string, res string, protected bool) *httptest.Server {
	if ct == "" {
		ct = "text/html"
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if protected {
			u, p, ok := r.BasicAuth()
			if !ok || (u != "test" && p != "test") {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		w.Header().Set("Content-Type", ct)
		w.Write([]byte(res))
	}))
}
