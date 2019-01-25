package nest

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const cameraResponse = `{
    "name": "Den (dbec)",
    "software_version": "1.0.2",
    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnjnpDtxGHsWA",
    "device_id": "kphN5lNgHsDtoJkfKnDURMABSChmjsFcjoGuBimqasah81-lE93RiA",
    "structure_id": "xYylA-lypQl5FHuJj2pY_JU3k-aKvEOT3oUEV_Nu8u85w_hO-s4xRg",
    "is_online": true,
    "is_streaming": false,
    "is_audio_input_enabled": true,
    "last_is_online_change": "2015-08-03T04:05:06.637Z",
    "is_video_history_enabled": true,
    "is_public_share_enabled": true,
    "snapshot_url": "https://developer.nest.com/simulator/api/v1/nest/devices/camera/snapshot",
    "activity_zones": [
        {
            "name": "sample activity zone",
            "id": 0
        }
    ],
    "where_name": "Den",
    "name_long": "Den Camera (dbec)",
    "web_url": "https://home.nest.com/cameras/CjZrcGhONWxOZ0hzRHRvSmtmS25EVVJNQUJTQ2htanNGY2pvR3VCaW1xYXNhaDgxLWxFOTNSaUESFl8xWk9fZTBhTGUtYjl4NVM0UV96Q0EaNnBJRUVWOGxQMVY0NGViNXY2N3Jlc21HTUswa2tTempJOXBuTl9Ia0J1T0VPOU0zZU5MdkV5UQ?auth=4796phrhYE9YRCSmJ841DqKAAxvHd9C75WhJEIEwTweFSXUVcny00sxVRPcs5vMqMu9cjrQs6Kn3jnVzy3BMcJgt6-V10pviEeXjeZo1Ts9ViaqESc6VtCtPKHW5jEn30nge3stfG70YIjmy9HjGeScXfwtBzf4Wz3j2C_2MC_yf6sLx6rKFMSiuXEGao7h6LHkywc35ybwu0A",
    "app_url": "nestmobile://cameras/CjZrcGhONWxOZ0hzRHRvSmtmS25EVVJNQUJTQ2htanNGY2pvR3VCaW1xYXNhaDgxLWxFOTNSaUESFl8xWk9fZTBhTGUtYjl4NVM0UV96Q0EaNnBJRUVWOGxQMVY0NGViNXY2N3Jlc21HTUswa2tTempJOXBuTl9Ia0J1T0VPOU0zZU5MdkV5UQ?auth=4796phrhYE9YRCSmJ841DqKAAxvHd9C75WhJEIEwTweFSXUVcny00sxVRPcs5vMqMu9cjrQs6Kn3jnVzy3BMcJgt6-V10pviEeXjeZo1Ts9ViaqESc6VtCtPKHW5jEn30nge3stfG70YIjmy9HjGeScXfwtBzf4Wz3j2C_2MC_yf6sLx6rKFMSiuXEGao7h6LHkywc35ybwu0A"
}`

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

func Test_StreamCameraDevice(t *testing.T) {
	cl := newTestClient("event: 123\ndata: 456\n", http.StatusOK)
	ts := NewCameraService(cl)
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
