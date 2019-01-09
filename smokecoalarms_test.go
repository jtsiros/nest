package nest

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var smokecoalarmResponse string

func init() {
	smokecoalarmResponse = readFileContents("testdata/smokecoalarm.json")
}

func Test_GetSmokeCoAlarm(t *testing.T) {
	c := newTestClient(smokecoalarmResponse, http.StatusOK)
	s := NewSmokeCoAlarmService(c)
	th, err := s.Get("2y2eUoaaXxBij4O1rxoiGfVfehTNCJA_")
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "2y2eUoaaXxBij4O1rxoiGfVfehTNCJA_", th.DeviceID, "deviceIDs should match")
}
