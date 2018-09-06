package nest

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/jtsiros/nest/device"
	"github.com/jtsiros/nest/types/temperature"
)

type hvacMode string

const (
	// Heat mode
	Heat hvacMode = "heat"
	// Cool mode
	Cool hvacMode = "cool"
	// HeatCool mode
	HeatCool hvacMode = "heat-cool"
	// Eco mode
	Eco hvacMode = "eco"
	// Off mode
	Off hvacMode = "off"
)

type tempScale string

const (
	// F represents Farenheight
	F tempScale = "F"
	// C represents celcius
	C tempScale = "C"
)

type values map[string]interface{}

// ThermostatService interacts with Thermostat devices to control and read device data.
type ThermostatService service

// NewThermostatService creates a new service to interact with Thermostats.
func NewThermostatService(client *Client) *ThermostatService {
	rel := &url.URL{Path: "/devices/thermostats"}
	u := client.baseURL.ResolveReference(rel)

	return &ThermostatService{
		client: client,
		apiURL: u,
	}
}

// SetTargetTemperature changes the target temperature on the Thermostat.
// See https://developers.nest.com/guides/thermostat-guide#target_temperature
//
func (svc *ThermostatService) SetTargetTemperature(deviceid string, scale string, target float64) error {
	ttKey := fmt.Sprintf("target_temperature_%s", strings.ToLower(scale))
	return svc.requestWithValues(svc.apiURL.String()+deviceid, values{ttKey: target})
}

// SetTargetTemperatureRange changes the target temperature on the Thermostat with a given range.
// See https://developers.nest.com/guides/thermostat-guide#
// target_temperature_low(f|c)
// target_temperature_high(f|c)
//
func (svc *ThermostatService) SetTargetTemperatureRange(deviceid string, scale string, low temperature.Temperature, high temperature.Temperature) error {
	scale = strings.ToLower(scale)
	values := map[string]interface{}{}

	if low == nil && high == nil {
		return errors.New("either low or high target must be set")
	}
	if low != nil && high != nil && *low >= *high {
		return errors.New("low value must be less than high value")
	}

	if low != nil {
		lowKey := fmt.Sprintf("target_temperature_low_%s", scale)
		values[lowKey] = *low
	}
	if high != nil {
		highKey := fmt.Sprintf("target_temperature_high_%s", scale)
		values[highKey] = *high
	}

	return svc.requestWithValues(svc.apiURL.String()+deviceid, values)
}

// SetHVACMode sets thermostat to the given mode. Current modes supported: (heat, cool, heat-cool, eco, off)
// Indicates HVAC system heating/cooling modes, like Heat•Cool for systems with heating and cooling capacity,
// or Eco Temperatures for energy savings.
//
// See https://developers.nest.com/reference/api-thermostat#hvac_mode
//
func (svc *ThermostatService) SetHVACMode(deviceid string, state hvacMode) error {
	return svc.requestWithValues(svc.apiURL.String()+deviceid, values{"hvac_mode": state})
}

// SetFanTimerDuration specifies the length of time (in minutes) that the fan is set to run.
// See https://developers.nest.com/reference/api-thermostat#fan_timer_duration
//
func (svc *ThermostatService) SetFanTimerDuration(deviceid string, duration int) error {
	if duration%15 != 0 {
		return errors.New("duration must be a multiple of 15")
	}
	return svc.requestWithValues(svc.apiURL.String()+deviceid, values{"fan_timer_duration": duration})
}

// SetFanTimerActive specifies the length of time (in minutes) that the fan is set to run.
// See https://developers.nest.com/reference/api-thermostat#fan_timer_duration
//
func (svc *ThermostatService) SetFanTimerActive(deviceid string, active bool) error {
	return svc.requestWithValues(svc.apiURL.String()+deviceid, values{"fan_timer_active": active})
}

// SetLabel sets a custom label for a thermostat.
// See https://developers.nest.com/reference/api-thermostat#label
func (svc *ThermostatService) SetLabel(deviceid string, label string) error {
	return svc.requestWithValues(svc.apiURL.String()+deviceid, values{"label": label})
}

// SetTemperatureScale sets the temperature scale display to F or C.
func (svc *ThermostatService) SetTemperatureScale(deviceid string, scale tempScale) error {
	return svc.requestWithValues(svc.apiURL.String()+deviceid, values{"temperature_scale": scale})
}

// Get fetches an updated thermostat object given a deviceID.
// https://developers.nest.com/guides/api/thermostat-guide
// See Thermostat Identifiers
//
func (svc *ThermostatService) Get(deviceid string) (*device.Thermostat, error) {
	var thermostat device.Thermostat
	err := svc.client.getDevice(deviceid, svc.apiURL.String(), &thermostat)
	return &thermostat, err
}

func (svc *ThermostatService) requestWithValues(path string, values map[string]interface{}) error {
	req, err := svc.client.newRequest("PUT", path, values)
	if err != nil {
		return err
	}

	_, err = svc.client.do(req, nil)
	return err
}
