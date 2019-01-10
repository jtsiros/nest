package nest

import (
	"net/http"
	"testing"

	"github.com/jtsiros/nest/device"
	"github.com/stretchr/testify/assert"
)

var thermostatResponse string

func init() {
	thermostatResponse = readFileContents("testdata/thermostat.json")
}

func Test_GetThermostat(t *testing.T) {
	c := newTestClient(thermostatResponse, http.StatusOK)
	s := NewThermostatService(c)

	th, err := s.Get("JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_")
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
	assert.Equal(t, "JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_", th.DeviceID, "deviceIDs should match")
}

func Test_SetTargetTemp(t *testing.T) {

	tt := []struct {
		target float64
		s      *ThermostatService
		d      *device.Thermostat
		err    string
	}{
		{76.0, NewThermostatService(newTestClient("76", http.StatusOK)), &device.Thermostat{DeviceID: "12345abcd", TemperatureScale: "F"}, ""},
		{91.0, NewThermostatService(newTestClient("{\"message\": \"Temperature F value is too high: 91\"}", http.StatusBadRequest)), &device.Thermostat{DeviceID: "12345abcd", TemperatureScale: "F"}, "Temperature F value is too high: 91"},
		{10.0, NewThermostatService(newTestClient("{\"message\": \"Temperature F value is too low: 10\"}", http.StatusBadRequest)), &device.Thermostat{DeviceID: "12345abcd", TemperatureScale: "F"}, "Temperature F value is too low: 10"},
	}

	for _, tc := range tt {
		err := tc.s.SetTargetTemperature(tc.d.DeviceID, tempScale(tc.d.TemperatureScale), tc.target)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}

func Test_SetTargetTempRange(t *testing.T) {

	tt := []struct {
		low  float64
		high float64
		s    *ThermostatService
		d    *device.Thermostat
		err  string
	}{
		{76.0, 80.0, NewThermostatService(newTestClient("76", http.StatusOK)), &device.Thermostat{DeviceID: "12345abcd", TemperatureScale: "F"}, ""},
		{80.0, 80.0, NewThermostatService(newTestClient("", http.StatusOK)), &device.Thermostat{DeviceID: "12345abcd", TemperatureScale: "F"}, ""},
		{80.0, 76.0, NewThermostatService(newTestClient("", http.StatusBadRequest)), &device.Thermostat{DeviceID: "12345abcd", TemperatureScale: "F"}, "low value must be less than or equal to high value"},
		{0.0, 0.0, NewThermostatService(newTestClient("", http.StatusBadRequest)), &device.Thermostat{DeviceID: "12345abcd", TemperatureScale: "F"}, "either low or high target must be set above 0"},
	}

	for _, tc := range tt {
		err := tc.s.SetTargetTemperatureRange(tc.d.DeviceID, tempScale(tc.d.TemperatureScale), tc.low, tc.high)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}

func Test_SetHVACMode(t *testing.T) {

	tt := []struct {
		deviceID string
		s        *ThermostatService
		mode     hvacMode
		err      string
	}{
		{"123", NewThermostatService(newTestClient("heat", http.StatusOK)), Heat, ""},
		{"123", NewThermostatService(newTestClient("cool", http.StatusOK)), Cool, ""},
		{"123", NewThermostatService(newTestClient("heat-cool", http.StatusOK)), HeatCool, ""},
		{"123", NewThermostatService(newTestClient("eco", http.StatusOK)), Eco, ""},
		{"123", NewThermostatService(newTestClient("off", http.StatusOK)), Eco, ""},
		{"123", NewThermostatService(newTestClient("{\"message\":\"Invalid HVAC mode: blah\"}", http.StatusBadRequest)), hvacMode("blah"), "Invalid HVAC mode: blah"},
		{"456", NewThermostatService(newTestClient("{\"message\":\"Invalid thermostat id: 456\"}", http.StatusBadRequest)), Cool, "Invalid thermostat id: 456"},
	}

	for _, tc := range tt {
		err := tc.s.SetHVACMode(tc.deviceID, tc.mode)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}

func Test_SetFanTimerDuration(t *testing.T) {
	tt := []struct {
		deviceID string
		s        *ThermostatService
		duration int
		err      string
	}{
		{"123", NewThermostatService(newTestClient("{\"message\": \"Cannot set fan_timer_duration to the selected value. See API reference for allowed values.\"}", http.StatusBadRequest)), 1500, "Cannot set fan_timer_duration to the selected value. See API reference for allowed values."},
		{"123", NewThermostatService(newTestClient("", http.StatusBadRequest)), 10, "duration must be a multiple of 15"},
		{"123", NewThermostatService(newTestClient("", http.StatusOK)), 15, ""},
	}

	for _, tc := range tt {
		err := tc.s.SetFanTimerDuration(tc.deviceID, tc.duration)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}

func Test_GetFanTimerActive(t *testing.T) {
	tt := []struct {
		deviceID string
		s        *ThermostatService
		active   bool
		err      string
	}{
		{"123", NewThermostatService(newTestClient("true", http.StatusBadRequest)), true, ""},
		{"123", NewThermostatService(newTestClient("false", http.StatusBadRequest)), false, ""},
	}

	for _, tc := range tt {
		err := tc.s.GetFanTimerActive(tc.deviceID)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}

func Test_SetLabel(t *testing.T) {
	tt := []struct {
		deviceID string
		s        *ThermostatService
		label    string
		err      string
	}{
		{"123", NewThermostatService(newTestClient("{\"message\":\"Label must be less than 256 characters\"}", http.StatusBadRequest)), "extra long label over 256 characters", "Label must be less than 256 characters"},
		{"123", NewThermostatService(newTestClient("", http.StatusOK)), "Simple", ""},
	}

	for _, tc := range tt {
		err := tc.s.SetLabel(tc.deviceID, tc.label)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}

func Test_SetTempScale(t *testing.T) {
	tt := []struct {
		deviceID string
		s        *ThermostatService
		scale    tempScale
		err      string
	}{
		{"123", NewThermostatService(newTestClient("{\"message\":\"Unspecified error\"}", http.StatusBadRequest)), tempScale("D"), "Unspecified error"},
		{"123", NewThermostatService(newTestClient("F", http.StatusOK)), F, ""},
		{"123", NewThermostatService(newTestClient("C", http.StatusOK)), C, ""},
	}

	for _, tc := range tt {
		err := tc.s.SetTemperatureScale(tc.deviceID, tc.scale)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}
