package nest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jtsiros/nest/config"
	"github.com/stretchr/testify/assert"
)

func Test_ListOfDevices(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, APIResponse)
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
