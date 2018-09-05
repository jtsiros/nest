package nest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jtsiros/nest/device"

	"github.com/jtsiros/nest/config"
	"github.com/stretchr/testify/assert"
)

var apiResponse string
var deviceResponse string

func init() {
	apiResponse = readFileContents("testdata/apiResponse.json")
	deviceResponse = readFileContents("testdata/device.json")
}

func readFileContents(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Test_ListOfDevices(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, apiResponse)
	}))

	api, err := NewClient(config.Config{APIURL: ts.URL}, ts.Client())
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	devices, err := api.Devices()
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
	assert.True(t, len(devices.Thermostats) > 0, "should return at least one thermostat")
}

func Test_GetThermostat(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, deviceResponse)
	}))

	api, err := NewClient(config.Config{APIURL: ts.URL}, ts.Client())
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	th, err := api.Thermostats.Get("JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_")
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_", th.DeviceID, "deviceIDs should match")
}

func Test_SetTargetTemp(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	}))

	api, err := NewClient(config.Config{APIURL: ts.URL}, ts.Client())
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	d := device.Thermostat{
		DeviceID:         "12345abcd",
		TemperatureScale: "F",
	}
	err = api.Thermostats.SetTargetTemperature(d.DeviceID, d.TemperatureScale, 76.0)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Nil(t, err)
}
