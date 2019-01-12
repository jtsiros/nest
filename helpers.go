package nest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type response struct {
	Status int
	Body   string
}

func newTestServer(responseMap map[string]*response) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if res, ok := responseMap[r.URL.String()]; ok {
			w.WriteHeader(res.Status)
			fmt.Fprintln(w, res.Body)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}))
}

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
