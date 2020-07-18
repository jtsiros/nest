package nest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func newTestClient(response string, code int) *Client {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		fmt.Fprintln(w, response)
	}))
	return newTestClientWithServer(ts)
}

func newTestClientWithServer(ts *httptest.Server) *Client {
	url, _ := url.Parse(ts.URL)
	c := &Client{
		baseURL:    url,
		httpClient: ts.Client(),
	}
	return c
}
