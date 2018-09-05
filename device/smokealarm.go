package device

import "time"

// SmokeAlarm is a Nest SmokeAlarm designed to detect smoke and carbon monoxide
// See https://developers.nest.com/documentation/cloud/smoke-co-guide
type SmokeAlarm struct {
	Locale             string    `json:"locale,omitempty"`
	StructureID        string    `json:"structure_id,omitempty"`
	SoftwareVersion    string    `json:"software_version,omitempty"`
	WhereID            string    `json:"where_id,omitempty"`
	DeviceID           string    `json:"device_id,omitempty"`
	WhereName          string    `json:"where_name,omitempty"`
	Name               string    `json:"name,omitempty"`
	NameLong           string    `json:"name_long,omitempty"`
	IsOnline           bool      `json:"is_online,omitempty"`
	LastConnection     time.Time `json:"last_connection,omitempty"`
	BatteryHealth      string    `json:"battery_health,omitempty"`
	CoAlarmState       string    `json:"co_alarm_state,omitempty"`
	SmokeAlarmState    string    `json:"smoke_alarm_state,omitempty"`
	UIColorState       string    `json:"ui_color_state,omitempty"`
	IsManualTestActive bool      `json:"is_manual_test_active,omitempty"`
}
