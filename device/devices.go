package device

// Devices represent a collection of Nest devices
// https://developers.nest.com/documentation/cloud/api-overview
//
type Devices struct {
	Thermostats   map[string]*Thermostat `json:"thermostats,omitempty"`
	SmokeCoAlarms map[string]*SmokeAlarm `json:"smoke_co_alarms,omitempty"`
	Cameras       map[string]*Camera     `json:"cameras,omitempty"`
}

// Len returns a count of the number of devices from the Nest API.
func (d Devices) Len() int {
	return len(d.Cameras) + len(d.SmokeCoAlarms) + len(d.Thermostats)
}
