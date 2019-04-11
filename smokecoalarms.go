package nest

import (
	"fmt"
	"net/url"

	"github.com/jtsiros/nest/config"
	"github.com/jtsiros/nest/device"
)

// SmokeCoAlarmService interacts with Smoke+CO Alarm devices to read device data.
type SmokeCoAlarmService service

// NewSmokeCoAlarmService creates a new service to interact with Smoke+CO alarms.
func NewSmokeCoAlarmService(client *Client) *SmokeCoAlarmService {
	u := &url.URL{Path: "/devices/smoke_co_alarms"}

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

// Stream opens an event stream to monitor changes on the smokecoalarm.
// https://developers.nest.com/guides/api/rest-streaming-guide
// https://developers.nest.com/reference/api-smoke-co-alarm
//
func (svc *SmokeCoAlarmService) Stream(deviceID string) (*Stream, error) {
	rel := &url.URL{Path: fmt.Sprintf("%s/%s", svc.apiURL.String(), deviceID)}
	return NewStream(&config.Config{
		APIURL: svc.client.baseURL.ResolveReference(rel).String(),
	}, svc.client.httpClient)
}
