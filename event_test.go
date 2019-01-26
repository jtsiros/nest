package nest

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jtsiros/nest/config"
	"github.com/jtsiros/nest/device"
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

func Test_getEvent(t *testing.T) {
	thermostat := createHandler("event: thermostats\ndata: {\"path\":\"/devices/thermostats/1234\",\"data\":{\"device_id\":\"1234\"}}\n")
	smokeCoAlarm := createHandler("event: smoke_co_alarms\ndata: {\"path\":\"/devices/smoke_co_alarms/1234\",\"data\":{\"device_id\":\"1234\"}}\n")
	camera := createHandler("event: cameras\ndata: {\"path\":\"/devices/cameras/1234\",\"data\":{\"device_id\":\"1234\"}}\n")
	keepAlive := createHandler("event: keep-alive\ndata: \n")
	eventError := createHandler("event: error\ndata: \n")

	req := httptest.NewRequest("GET", "http://localhost/", nil)
	tt := []struct {
		req        func(http.ResponseWriter, *http.Request)
		rec        *httptest.ResponseRecorder
		deviceID   string
		deviceType EventsType
		eventType  reflect.Type
	}{
		{thermostat, httptest.NewRecorder(), "1234", Thermostats, reflect.TypeOf(&device.Thermostat{})},
		{smokeCoAlarm, httptest.NewRecorder(), "1234", SmokeCoAlarms, reflect.TypeOf(&device.SmokeAlarm{})},
		{camera, httptest.NewRecorder(), "1234", Cameras, reflect.TypeOf(&device.Camera{})},
		{keepAlive, httptest.NewRecorder(), "", KeepAlive, nil},
		{eventError, httptest.NewRecorder(), "", EventError, nil},
	}

	for _, tc := range tt {
		tc.req(tc.rec, req)
		resp := tc.rec.Result()

		events := make(chan Event)
		go readEvents(events, resp)
		event := <-events
		et, deviceID, device, _ := event.GetEvent(false)
		assert.Equal(t, tc.deviceID, deviceID)
		assert.Equal(t, tc.deviceType, et)
		assert.Equal(t, tc.eventType, reflect.TypeOf(device))
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
