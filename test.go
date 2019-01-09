package nest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var APIResponse string

func init() {
	APIResponse = readFileContents("testdata/apiResponse.json")
}

func readFileContents(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func newTestClient(response string, code int) *Client {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		fmt.Fprintln(w, response)
	}))

	url, _ := url.Parse(ts.URL)
	c := &Client{
		baseURL:    url,
		httpClient: ts.Client(),
	}
	return c
}
