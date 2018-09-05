package nest

import (
	"net/url"

	"github.com/jtsiros/nest/device"
)

// CameraService provides read and control for Camera devices. Most of the device
// calls are read-only, thus a majority of the time, a call to Get() will fetch
// all of the appropriate attributes although convenience methods for writing
// certain values are provided.
type CameraService service

// NewCameraService creates a new service.
func NewCameraService(client *Client) *CameraService {
	rel := &url.URL{Path: "/devices/cameras"}
	u := client.baseURL.ResolveReference(rel)

	return &CameraService{
		client: client,
		apiURL: u,
	}
}

// Get fetches an updated camera object given a device id.
// https://developers.nest.com/reference/api-camera
//
func (svc *CameraService) Get(deviceid string) (*device.Camera, error) {
	var camera device.Camera
	err := svc.client.getDevice(deviceid, svc.apiURL.String(), &camera)
	return &camera, err
}
