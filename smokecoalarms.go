package nest

import (
	"net/url"

	"github.com/jtsiros/nest/device"
)

// SmokeCoAlarmService interacts with Smoke+CO Alarm devices to read device data.
type SmokeCoAlarmService service

// NewSmokeCoAlarmService creates a new service to interact with Smoke+CO alarms.
func NewSmokeCoAlarmService(client *Client) *SmokeCoAlarmService {
	rel := &url.URL{Path: "/devices/smoke_co_alarms"}
	u := client.baseURL.ResolveReference(rel)

	return &SmokeCoAlarmService{
		client: client,
		apiURL: u,
	}
}

// Get fetches an updated smokecoalarm object given a device id.
// https://developers.nest.com/reference/api-smoke-co-alarm
//
func (svc *SmokeCoAlarmService) Get(deviceid string) (*device.SmokeAlarm, error) {
	var smokeCoAlarm device.SmokeAlarm
	err := svc.client.getDevice(deviceid, svc.apiURL.String(), &smokeCoAlarm)
	return &smokeCoAlarm, err
}
