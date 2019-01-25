package nest

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jtsiros/nest/device"
	"github.com/stretchr/testify/assert"
)

const thermostatResponse = `{
    "humidity": 50,
    "locale": "en-US",
    "temperature_scale": "F",
    "is_using_emergency_heat": false,
    "has_fan": true,
    "software_version": "5.6.1",
    "has_leaf": true,
    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkq85vhsG-Xhg",
    "device_id": "JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_",
    "name": "Family Room (ADCD)",
    "can_heat": true,
    "can_cool": true,
    "target_temperature_c": 20,
    "target_temperature_f": 68,
    "target_temperature_high_c": 26,
    "target_temperature_high_f": 79,
    "target_temperature_low_c": 19,
    "target_temperature_low_f": 66,
    "ambient_temperature_c": 21,
    "ambient_temperature_f": 70,
    "away_temperature_high_c": 24,
    "away_temperature_high_f": 76,
    "away_temperature_low_c": 12.5,
    "away_temperature_low_f": 55,
    "eco_temperature_high_c": 24,
    "eco_temperature_high_f": 76,
    "eco_temperature_low_c": 12.5,
    "eco_temperature_low_f": 55,
    "is_locked": false,
    "locked_temp_min_c": 20,
    "locked_temp_min_f": 68,
    "locked_temp_max_c": 22,
    "locked_temp_max_f": 72,
    "sunlight_correction_active": false,
    "sunlight_correction_enabled": true,
    "structure_id": "xYylA-lypQl5FHuJj2pY_JU3k-aKvEOT3oUEV_Nu8u85w_hO-s4xRg",
    "fan_timer_active": false,
    "fan_timer_timeout": "1970-01-01T00:00:00.000Z",
    "fan_timer_duration": 15,
    "previous_hvac_mode": "",
    "hvac_mode": "heat",
    "time_to_target": "~0",
    "time_to_target_training": "ready",
    "where_name": "Family Room",
    "label": "ADCD",
    "name_long": "Family Room Thermostat (ADCD)",
    "is_online": true,
    "hvac_state": "off"
}`

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

	d := &device.Thermostat{
		DeviceID:         "12345abcd",
		TemperatureScale: "F",
	}

	tt := []struct {
		target float64
		s      *ThermostatService
		d      *device.Thermostat
		err    string
	}{
		{76.0, NewThermostatService(newTestClient("76", http.StatusOK)), d, ""},
		{91.0, NewThermostatService(newTestClient("{\"message\": \"Temperature F value is too high: 91\"}", http.StatusBadRequest)), d, "Temperature F value is too high: 91"},
		{10.0, NewThermostatService(newTestClient("{\"message\": \"Temperature F value is too low: 10\"}", http.StatusBadRequest)), d, "Temperature F value is too low: 10"},
	}

	for _, tc := range tt {
		err := tc.s.SetTargetTemperature(tc.d.DeviceID, tempScale(tc.d.TemperatureScale), tc.target)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
		}
	}
}

func Test_SetTargetTempRange(t *testing.T) {

	d := &device.Thermostat{
		DeviceID:         "12345abcd",
		TemperatureScale: "F",
	}

	tt := []struct {
		low  float64
		high float64
		s    *ThermostatService
		d    *device.Thermostat
		err  string
	}{
		{76.0, 80.0, NewThermostatService(newTestClient("76", http.StatusOK)), d, ""},
		{80.0, 80.0, NewThermostatService(newTestClient("", http.StatusOK)), d, ""},
		{80.0, 76.0, NewThermostatService(newTestClient("", http.StatusBadRequest)), d, "low value must be less than or equal to high value"},
		{0.0, 0.0, NewThermostatService(newTestClient("", http.StatusBadRequest)), d, "either low or high target must be set above 0"},
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

func Test_requestWithValues(t *testing.T) {

	tt := []struct {
		method string
		path   string
		values map[string]interface{}
		s      *ThermostatService
		err    error
	}{
		{http.MethodGet, "/test", nil, NewThermostatService(newTestClient("", http.StatusOK)), nil},
		{"()", "/test", nil, NewThermostatService(newTestClient("", http.StatusOK)), errors.New("net/http: invalid method \"()\"")},
	}

	for _, tc := range tt {
		err := tc.s.requestWithValues(tc.method, tc.path, tc.values)
		if tc.err != nil {
			isNotNil := assert.NotNil(t, err)
			if isNotNil {
				assert.Equal(t, tc.err.Error(), err.Error())
			}
		}
	}
}

func Test_StreamThermostatDevice(t *testing.T) {
	cl := newTestClient("event: 123\ndata: 456\n", http.StatusOK)
	ts := NewThermostatService(cl)
	s, err := ts.Stream("12345")
	if err != nil {
		t.Fatal(err)
	}

	c, err := s.Open()
	if err != nil {
		t.Fatal(err)
	}
	event := <-c
	assert.Equal(t, []byte("123"), event.name)
	assert.Equal(t, []byte("456"), event.data)
}
