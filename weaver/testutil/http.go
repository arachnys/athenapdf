package testutil

import (
	"net/http"
	"net/http/httptest"
)

func MockHTTPServer(ct string, res string) *httptest.Server {
	if ct == "" {
		ct = "text/html"
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ct)
		w.Write([]byte(res))
	}))
}
