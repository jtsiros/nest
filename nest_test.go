package nest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jtsiros/nest/config"
	"github.com/stretchr/testify/assert"
)

const apiResponse = `{
    "devices": {
        "thermostats": {
            "JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_": {
                "humidity": 50,
                "locale": "en-US",
                "temperature_scale": "F",
                "is_using_emergency_heat": false,
                "has_fan": true,
                "software_version": "5.6.1",
                "has_leaf": true,
                "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkq85vhsG-Xhg",
                "device_id": "JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_",
                "name": "Family Room (ADCD)",
                "can_heat": true,
                "can_cool": true,
                "target_temperature_c": 20,
                "target_temperature_f": 68,
                "target_temperature_high_c": 26,
                "target_temperature_high_f": 79,
                "target_temperature_low_c": 19,
                "target_temperature_low_f": 66,
                "ambient_temperature_c": 21,
                "ambient_temperature_f": 70,
                "away_temperature_high_c": 24,
                "away_temperature_high_f": 76,
                "away_temperature_low_c": 12.5,
                "away_temperature_low_f": 55,
                "eco_temperature_high_c": 24,
                "eco_temperature_high_f": 76,
                "eco_temperature_low_c": 12.5,
                "eco_temperature_low_f": 55,
                "is_locked": false,
                "locked_temp_min_c": 20,
                "locked_temp_min_f": 68,
                "locked_temp_max_c": 22,
                "locked_temp_max_f": 72,
                "sunlight_correction_active": false,
                "sunlight_correction_enabled": true,
                "structure_id": "xYylA-lypQl5FHuJj2pY_JU3k-aKvEOT3oUEV_Nu8u85w_hO-s4xRg",
                "fan_timer_active": false,
                "fan_timer_timeout": "1970-01-01T00:00:00.000Z",
                "fan_timer_duration": 15,
                "previous_hvac_mode": "",
                "hvac_mode": "heat",
                "time_to_target": "~0",
                "time_to_target_training": "ready",
                "where_name": "Family Room",
                "label": "ADCD",
                "name_long": "Family Room Thermostat (ADCD)",
                "is_online": true,
                "hvac_state": "off"
            },
            "JP2FgJUZqqBppN16wdLGxvVfehTNCJA_": {
                "humidity": 50,
                "locale": "en-US",
                "temperature_scale": "F",
                "is_using_emergency_heat": false,
                "has_fan": true,
                "software_version": "5.6.1",
                "has_leaf": true,
                "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnjnpDtxGHsWA",
                "device_id": "JP2FgJUZqqBppN16wdLGxvVfehTNCJA_",
                "name": "Den (3A8E)",
                "can_heat": true,
                "can_cool": true,
                "target_temperature_c": 20,
                "target_temperature_f": 68,
                "target_temperature_high_c": 26,
                "target_temperature_high_f": 79,
                "target_temperature_low_c": 19,
                "target_temperature_low_f": 66,
                "ambient_temperature_c": 21,
                "ambient_temperature_f": 70,
                "away_temperature_high_c": 24,
                "away_temperature_high_f": 76,
                "away_temperature_low_c": 12.5,
                "away_temperature_low_f": 55,
                "eco_temperature_high_c": 24,
                "eco_temperature_high_f": 76,
                "eco_temperature_low_c": 12.5,
                "eco_temperature_low_f": 55,
                "is_locked": false,
                "locked_temp_min_c": 20,
                "locked_temp_min_f": 68,
                "locked_temp_max_c": 22,
                "locked_temp_max_f": 72,
                "sunlight_correction_active": false,
                "sunlight_correction_enabled": true,
                "structure_id": "xYylA-lypQl5FHuJj2pY_JU3k-aKvEOT3oUEV_Nu8u85w_hO-s4xRg",
                "fan_timer_active": false,
                "fan_timer_timeout": "1970-01-01T00:00:00.000Z",
                "fan_timer_duration": 15,
                "previous_hvac_mode": "",
                "hvac_mode": "heat-cool",
                "time_to_target": "~0",
                "time_to_target_training": "ready",
                "where_name": "Den",
                "label": "3A8E",
                "name_long": "Den Thermostat (3A8E)",
                "is_online": true,
                "hvac_state": "off"
            }
        },
        "smoke_co_alarms": {
            "2y2eUoaaXxBij4O1rxoiGfVfehTNCJA_": {
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
            }
        },
        "cameras": {
            "kphN5lNgHsDtoJkfKnDURMABSChmjsFcjoGuBimqasah81-lE93RiA": {
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
            }
        }
    },
    "structures": {
        "xYylA-lypQl5FHuJj2pY_JU3k-aKvEOT3oUEV_Nu8u85w_hO-s4xRg": {
            "smoke_co_alarms": [
                "2y2eUoaaXxBij4O1rxoiGfVfehTNCJA_"
            ],
            "name": "Home 1",
            "country_code": "US",
            "time_zone": "America/Los_Angeles",
            "away": "home",
            "thermostats": [
                "JP2FgJUZqqAXUBfYYWVUY_VfehTNCJA_",
                "JP2FgJUZqqBppN16wdLGxvVfehTNCJA_"
            ],
            "structure_id": "xYylA-lypQl5FHuJj2pY_JU3k-aKvEOT3oUEV_Nu8u85w_hO-s4xRg",
            "co_alarm_state": "ok",
            "smoke_alarm_state": "ok",
            "wwn_security_state": "ok",
            "wheres": {
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCn9-0TRb0IF3w": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCn9-0TRb0IF3w",
                    "name": "Attic"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCln0pq3mbnJQA": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCln0pq3mbnJQA",
                    "name": "Back Door"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClLABYeI0vSKw": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClLABYeI0vSKw",
                    "name": "Backyard"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkaWayJGJS8sg": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkaWayJGJS8sg",
                    "name": "Basement"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnsZSVLDWnX7g": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnsZSVLDWnX7g",
                    "name": "Bathroom"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk4wGauKzXtCg": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk4wGauKzXtCg",
                    "name": "Bedroom"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkkij9mkUrSaA": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkkij9mkUrSaA",
                    "name": "Deck"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnjnpDtxGHsWA": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnjnpDtxGHsWA",
                    "name": "Den"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCldOXy3o7rCsg": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCldOXy3o7rCsg",
                    "name": "Dining Room"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk_rjo2JmLyWg": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk_rjo2JmLyWg",
                    "name": "Downstairs"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCktu8wgUyWE2g": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCktu8wgUyWE2g",
                    "name": "Driveway"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClYU0jvnAXMbw": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClYU0jvnAXMbw",
                    "name": "Entryway"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkq85vhsG-Xhg": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkq85vhsG-Xhg",
                    "name": "Family Room"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCltfRQf9yM8Fw": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCltfRQf9yM8Fw",
                    "name": "Front Door"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmID8EVd_hI4Q": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmID8EVd_hI4Q",
                    "name": "Front Yard"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmJ-ByghuMV-Q": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmJ-ByghuMV-Q",
                    "name": "Garage"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmMbsSCZk3OvQ": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmMbsSCZk3OvQ",
                    "name": "Guest House"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk2kXgI-jld4g": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk2kXgI-jld4g",
                    "name": "Guest Room"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClg5Knsky5T0A": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClg5Knsky5T0A",
                    "name": "Hallway"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCl5CKXHPZu1TA": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCl5CKXHPZu1TA",
                    "name": "Kids Room"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmoi76_UqAhzw": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCmoi76_UqAhzw",
                    "name": "Kitchen"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkINAIgxnJtnQ": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCkINAIgxnJtnQ",
                    "name": "Living Room"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnT33VUdTohNw": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnT33VUdTohNw",
                    "name": "Master Bedroom"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk1ukAK2RwEhA": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCk1ukAK2RwEhA",
                    "name": "Office"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClZ_Z93nrPpVA": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClZ_Z93nrPpVA",
                    "name": "Outside"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCl2vzr5faXIUg": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCl2vzr5faXIUg",
                    "name": "Patio"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClMLlIU3P_RCQ": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClMLlIU3P_RCQ",
                    "name": "Shed"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClkHHpkc6q35Q": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCClkHHpkc6q35Q",
                    "name": "Side Door"
                },
                "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnWU4jekHFwXg": {
                    "where_id": "osefycV5UZDoYWGlEtfvxzW9zlfCGfY3qyS_RoiaCCnWU4jekHFwXg",
                    "name": "Upstairs"
                }
            },
            "cameras": [
                "kphN5lNgHsDtoJkfKnDURMABSChmjsFcjoGuBimqasah81-lE93RiA"
            ]
        }
    },
    "metadata": {
        "access_token": "c.34rbzlNpss9h1nWO6UNUhSWSkNf61rbz6tAX8qnQhrBGzAUUiZPuWvQtj2utRysXYGFMekYQ56BTdS2irMYFOkAmnicjx9taoToYGdapBuhaxCJNOg97knt9gFVyZbmsHou0Q9VEKGcIuEvx",
        "client_version": 1,
        "user_id": "z.1.1.9kWKXfVeQUWQCfvHTsSTjFHMxTia08Rntt8UaFHwaiA="
    }
}
`

func Test_ListOfDevices(t *testing.T) {
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, apiResponse)
	}))

	tsErr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"message": "error msg"}`)
	}))

	tt := []struct {
		s   *httptest.Server
		err error
	}{
		{tsSuccess, nil},
		{tsErr, errors.New("error msg")},
	}

	for _, tc := range tt {
		api, err := NewClient(config.Config{APIURL: tc.s.URL}, tc.s.Client())
		devices, err := api.Devices()
		if tc.err != nil {
			if err == nil {
				t.Fatal("expected non nil Devices() response")
			}
			assert.Equal(t, tc.err.Error(), err.Error())
		} else {
			assert.True(t, len(devices.Thermostats) > 0, "should return at least one thermostat")
		}
	}
}

func Test_NewClientAPIUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, apiResponse)
	}))

	tt := []struct {
		url string
		err error
	}{
		{"http://localhost", nil},
		{"()://", errors.New("Not a properly formed API URL: parse \"()://\": first path segment in URL cannot contain colon")},
	}

	for _, tc := range tt {
		_, err := NewClient(config.Config{APIURL: tc.url}, ts.Client())
		if tc.err != nil {
			if tc.err.Error() != err.Error() {
				t.Fatalf("expected client error [%v], got [%v]", tc.err, err)
			}
		}
	}
}

func Test_NewRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, apiResponse)
	}))

	tt := []struct {
		method, path string
		body         interface{}
		expectedBody string
		err          error
	}{
		{"GET", "/test", map[string]string{"key": "value"}, "{\"key\":\"value\"}\n", nil},
		{"GET", "/test", map[string]string{}, "{}\n", nil},
		{"GET", "/test", nil, "", nil},
		{"POST", "/test", map[string]string{}, "{}\n", nil},
		{"()", "/test", nil, "", errors.New("net/http: invalid method \"()\"")},
	}

	c, _ := NewClient(config.Config{APIURL: ts.URL}, ts.Client())

	for _, tc := range tt {
		req, err := c.newRequest(tc.method, tc.path, tc.body)
		if tc.err != nil {
			if tc.err.Error() != err.Error() {
				t.Fatalf("expected client error [%v], got [%v]", tc.err, err)
			}
		} else {
			if req.Body != nil {
				b, _ := ioutil.ReadAll(req.Body)
				assert.Equal(t, tc.expectedBody, string(b))
			}
			assert.Equal(t, tc.method, req.Method)
		}
	}
}
