package nest

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const smokeCoAlarmResponse = `{
    "name": "Backyard (0305)",
    "locale": "en-US",
    "structure_id": "xYylA-lypQl5FHuJj2pY_JU3k-aKvEOT3oUEV_Nu8u85w_hO-s4xRg",
    "software_version": "1.0.2rc2",
    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClLABYeI0vSKw",
    "device_id": "2y2eUoaaXxBij4O1rxoiGfVfehTNCJA_",
    "where_name": "Backyard",
    "name_long": "Backyard Nest Protect (0305)",
    "is_online": true,
    "battery_health": "ok",
    "co_alarm_state": "ok",
    "smoke_alarm_state": "ok",
    "ui_color_state": "green",
    "is_manual_test_active": false,
    "last_manual_test_time": "2014-10-24T21:13:56.000Z"
}`

func Test_GetSmokeCoAlarm(t *testing.T) {
	c := newTestClient(smokeCoAlarmResponse, http.StatusOK)
	s := NewSmokeCoAlarmService(c)
	th, err := s.Get("2y2eUoaaXxBij4O1rxoiGfVfehTNCJA_")
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "2y2eUoaaXxBij4O1rxoiGfVfehTNCJA_", th.DeviceID, "deviceIDs should match")
}

func Test_StreamSmokeCoAlarmDevice(t *testing.T) {
	cl := newTestClient("event: 123\ndata: 456\n", http.StatusOK)
	ts := NewSmokeCoAlarmService(cl)
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
