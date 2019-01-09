package nest

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cameraResponse string

func init() {
	cameraResponse = readFileContents("testdata/camera.json")
}

func Test_GetCamera(t *testing.T) {
	c := newTestClient(cameraResponse, http.StatusOK)
	s := NewCameraService(c)
	th, err := s.Get("kphN5lNgHsDtoJkfKnDURMABSChmjsFcjoGuBimqasah81-lE93RiA")
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "kphN5lNgHsDtoJkfKnDURMABSChmjsFcjoGuBimqasah81-lE93RiA", th.DeviceID, "deviceIDs should match")
}
