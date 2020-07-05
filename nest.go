package nest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jtsiros/nest/config"
	"github.com/jtsiros/nest/device"
)

type service struct {
	client *Client
	apiURL *url.URL
}

// Client provides API functionality to Nest APIs
type Client struct {
	httpClient    *http.Client
	baseURL       *url.URL
	Thermostats   *ThermostatService
	SmokeCoAlarms *SmokeCoAlarmService
	Cameras       *CameraService
}

// Error represents an error from API call
type Error struct {
	Err      string `json:"error"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	Instance string `json:"instance"`
}

func (e Error) Error() string {
	return e.Message
}

// NewClient creates a new Nest API. Since the Nest API uses an authorization_code, there isn't an elegant way
// to prompt the API user to enter in an authorization code. This API assumes that the configured HTTP client
// is configured to handle OAuth2.
func NewClient(config config.Config, client *http.Client) (*Client, error) {
	url, err := url.Parse(config.APIURL)
	if err != nil {
		return nil, fmt.Errorf("not a properly formed API URL: %v", err)
	}
	c := &Client{
		baseURL:    url,
		httpClient: client,
	}

	c.Cameras = NewCameraService(c)
	c.Thermostats = NewThermostatService(c)
	c.SmokeCoAlarms = NewSmokeCoAlarmService(c)

	return c, nil
}

// Devices represent physical devices (Thermostats, Protects, and Cameras) within a structure
// https://developers.nest.com/documentation/cloud/architecture-overview
//
func (nest *Client) Devices() (*device.Devices, error) {
	req, err := nest.newRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}

	var d struct {
		Devices device.Devices
	}
	_, err = nest.do(req, &d)
	if err != nil {
		return nil, err
	}

	return &d.Devices, nil
}

// newRequest creates a well-formed http request given a method, relative path to base URL, and optional body.
// If a body is present, the assumed encoding is JSON. An authorization token is automatically added as a header.
func (nest *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := nest.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Add("Accept", "application/json")
	return req, err
}

func (nest *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := nest.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var err Error
		json.NewDecoder(resp.Body).Decode(&err)
		return nil, err
	}
	if v != nil {
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

func (nest *Client) getDevice(deviceid string, url string, device interface{}) error {
	req, err := nest.newRequest("GET", fmt.Sprintf("%s/%s", url, deviceid), nil)
	if err != nil {
		return err
	}

	_, err = nest.do(req, &device)
	return err
}
