package nest

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jtsiros/nest/config"
	"github.com/stretchr/testify/assert"
)

func Test_extract(t *testing.T) {

	tt := []struct {
		line     []byte
		prefix   []byte
		expected []byte
	}{
		{[]byte("data: this is my data\n"), []byte("data:"), []byte("this is my data")},
		{[]byte("event: update\n"), []byte("event:"), []byte("update")},
		{[]byte(""), []byte("event:"), nil},
	}

	for _, tc := range tt {
		res := extract(tc.line, tc.prefix)
		assert.Equal(t, tc.expected, res)
	}
}

func Test_readEvents(t *testing.T) {

	handlerSuccess := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "event: this is an event\ndata: this is data\n")
	}
	handlerEmpty := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://localhost/", nil)
	tt := []struct {
		req                        func(http.ResponseWriter, *http.Request)
		rec                        *httptest.ResponseRecorder
		expectedName, expectedData []byte
	}{
		{handlerSuccess, httptest.NewRecorder(), []byte("this is an event"), []byte("this is data")},
		{handlerEmpty, httptest.NewRecorder(), nil, nil},
	}

	for _, tc := range tt {
		tc.req(tc.rec, req)
		resp := tc.rec.Result()

		events := make(chan Event)
		go readEvents(events, resp)
		event := <-events

		assert.Equal(t, tc.expectedName, event.name)
		assert.Equal(t, tc.expectedData, event.data)
	}
}

func Test_createConnection(t *testing.T) {
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "name: event name\ndata: this is data")
	}))
	tsFailure := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"message":"error msg"}`)
	}))

	tt := []struct {
		expected string
		s        *httptest.Server
		err      error
	}{
		{"name: event name\ndata: this is data\n", tsSuccess, nil},
		{"", tsFailure, errors.New("expected status code 200, got 400")},
	}

	for _, tc := range tt {
		s, _ := NewStream(&config.Config{APIURL: tc.s.URL}, tc.s.Client())
		resp, err := s.createConnection()
		if tc.err != nil {
			if tc.err.Error() != err.Error() {
				t.Fatalf("expected err [%v] got [%v]\n", tc.err, err)
			}
		} else {
			b, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			assert.Equal(t, tc.expected, string(b))
		}
	}
}
