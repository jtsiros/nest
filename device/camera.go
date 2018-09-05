package device

import "time"

// Camera is a Nest Camera object. This is intended to work with all Nest Cam models
// and shows all Nest Cam data
// https://developers.nest.com/documentation/cloud/camera-guide
// https://developers.nest.com/documentation/cloud/api-camera
//
type Camera struct {
	DeviceID              string          `json:"device_id,omitempty"`
	SoftwareVersion       string          `json:"software_version,omitempty"`
	StructureID           string          `json:"structure_id,omitempty"`
	WhereID               string          `json:"where_id,omitempty"`
	WhereName             string          `json:"where_name,omitempty"`
	Name                  string          `json:"name,omitempty"`
	NameLong              string          `json:"name_long,omitempty"`
	IsOnline              bool            `json:"is_online,omitempty"`
	IsStreaming           bool            `json:"is_streaming,omitempty"`
	IsAudioInputEnabled   bool            `json:"is_audio_input_enabled,omitempty"`
	LastIsOnlineChange    time.Time       `json:"last_is_online_change,omitempty"`
	IsVideoHistoryEnabled bool            `json:"is_video_history_enabled,omitempty"`
	WebURL                string          `json:"web_url,omitempty"`
	AppURL                string          `json:"app_url,omitempty"`
	IsPublicShareEnabled  bool            `json:"is_public_share_enabled,omitempty"`
	ActivityZones         []*ActivityZone `json:"activity_zones,omitempty"`
	PublicShareURL        string          `json:"public_share_url,omitempty"`
	SnapshotURL           string          `json:"snapshot_url,omitempty"`
}

//ActivityZone represents an Activity Zone.
type ActivityZone struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
}
