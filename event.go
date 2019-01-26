package nest

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jtsiros/nest/config"
	"github.com/jtsiros/nest/device"
)

const (
	event = "event:"
	data  = "data:"
)

var (
	eventPfx = []byte(event)
	dataPfx  = []byte(data)
)

// EventsType describes the supported event types (usually based on device)
type EventsType string

// all the event types
const (
	Thermostats   EventsType = "thermostats"
	SmokeCoAlarms EventsType = "smoke_co_alarms"
	Cameras       EventsType = "cameras"
	KeepAlive     EventsType = "keep-alive"
	EventError    EventsType = "error"
)

// Event represents a response when changes occur in structure or device data.
type Event struct {
	name []byte
	data []byte
}

func (e Event) String() string {
	return fmt.Sprintf("Event name: %v, Data: %v", string(e.name), string(e.data))
}

type eventPath struct {
	Path string `json:"path"`
}

type thermostatData struct {
	Data *device.Thermostat `json:"data"`
}
type smokeCoAlarmData struct {
	Data *device.SmokeAlarm `json:"data"`
}
type cameraData struct {
	Data *device.Camera `json:"data"`
}

func prettyPrinter(eventType EventsType, v interface{}) (err error) {
	fmt.Printf("Event type: %v\n", eventType)
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// GetEvent returns the device data along with the type for type casting
// Returns the device type, device id, data, error
func (e Event) GetEvent(print bool) (EventsType, string, interface{}, error) {
	if string(e.name) == string(KeepAlive) {
		return KeepAlive, "", nil, nil
	}

	var eventPathBuf eventPath
	err := json.Unmarshal(e.data, &eventPathBuf)
	if err != nil {
		return EventError, "", nil, fmt.Errorf("path unmarshal error: %v", err)
	}

	// /devices/thermostats/<id>
	eventParsed := strings.Split(eventPathBuf.Path, "/")
	deviceID := eventParsed[3]
	eventType := EventsType(eventParsed[2])

	var deviceData interface{}
	switch eventType {
	case Thermostats:
		thermostatDataBuf := thermostatData{}
		err = json.Unmarshal(e.data, &thermostatDataBuf)
		deviceData = thermostatDataBuf.Data
	case SmokeCoAlarms:
		smokeCoAlarmDataBuf := smokeCoAlarmData{}
		err = json.Unmarshal(e.data, &smokeCoAlarmDataBuf)
		deviceData = smokeCoAlarmDataBuf.Data
	case Cameras:
		cameraDataBuf := cameraData{}
		err = json.Unmarshal(e.data, &cameraDataBuf)
		deviceData = cameraDataBuf.Data
	default:
		return EventError, "", nil, fmt.Errorf("unhandled device type: %v", eventType)
	}
	if err != nil {
		return EventError, "", nil, fmt.Errorf("data unmarshal error: %v", err)
	}

	if print {
		prettyPrinter(eventType, deviceData)
	}
	return eventType, deviceID, deviceData, nil
}

// Stream represents an open connection to the Nest APIs for device and structure changes.
// This will maintain an open socket for every stream connected to a device.
type Stream struct {
	client  *http.Client
	baseURL *url.URL
}

// NewStream returns a new stream given a configuration and http client objects.
func NewStream(cfg *config.Config, client *http.Client) (*Stream, error) {
	u, err := url.Parse(cfg.APIURL)
	if err != nil {
		return nil, err
	}
	return &Stream{
		client:  client,
		baseURL: u,
	}, nil
}

// Open opens a connection and streams events from the Nest API.
func (s Stream) Open() (chan Event, error) {
	events := make(chan Event)
	resp, err := s.createConnection()
	if err != nil {
		return nil, err
	}
	go readEvents(events, resp)
	return events, nil
}

// readEvents listens on the socket for events. Events are expected to contain
// event name, separated by newline, then event data.
// Once the event is assembled, the event is written to the events channel for consumption.
func readEvents(events chan<- Event, resp *http.Response) {
	reader := bufio.NewReader(resp.Body)
	event := Event{}
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Printf("error reading response: %v\n", err)
			break
		}
		if err == io.EOF {
			log.Printf("got EOF reading response: %v", err)
			break
		}

		switch {
		case bytes.HasPrefix(line, eventPfx):
			event.name = extract(line, eventPfx)
		case bytes.HasPrefix(line, dataPfx):
			event.data = extract(line, dataPfx)
			events <- event
			event = Event{}
		}
	}
	close(events)
}

func extract(line []byte, pfx []byte) []byte {
	if len(line) == 0 {
		return nil
	}
	pfxIdx := len(pfx) + 1
	endOfLineIdx := len(line) - 1
	return line[pfxIdx:endOfLineIdx]
}

// createConnection opens an event-stream to the Nest API to receive events from devices.
// This can be used to update the ambient temperature as it changes.
//
func (s Stream) createConnection() (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, s.baseURL.String(), nil)
	req.Header.Add("Accept", "text/event-stream")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to API: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	return resp, nil
}
