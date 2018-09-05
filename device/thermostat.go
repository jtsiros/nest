package device

import "time"

// Thermostat represents a Nest Thermostat that is internet-connected and accessible through the Nest API
type Thermostat struct {
	Humidity                  int       `json:"humidity,omitempty"`
	Locale                    string    `json:"locale,omitempty"`
	TemperatureScale          string    `json:"temperature_scale,omitempty"`
	IsUsingEmergencyHeat      bool      `json:"is_using_emergency_heat,omitempty"`
	HasFan                    bool      `json:"has_fan,omitempty,omitempty"`
	SoftwareVersion           string    `json:"software_version,omitempty"`
	HasLeaf                   bool      `json:"has_leaf,omitempty"`
	WhereID                   string    `json:"where_id,omitempty"`
	DeviceID                  string    `json:"device_id,omitempty"`
	Name                      string    `json:"name,omitempty"`
	CanHeat                   bool      `json:"can_heat,omitempty"`
	CanCool                   bool      `json:"can_cool,omitempty"`
	TargetTemperatureC        float64   `json:"target_temperature_c,omitempty"`
	TargetTemperatureF        int       `json:"target_temperature_f,omitempty"`
	TargetTemperatureHighC    float64   `json:"target_temperature_high_c,omitempty"`
	TargetTemperatureHighF    int       `json:"target_temperature_high_f,omitempty"`
	TargetTemperatureLowC     float64   `json:"target_temperature_low_c,omitempty"`
	TargetTemperatureLowF     int       `json:"target_temperature_low_f,omitempty"`
	AmbientTemperatureC       float64   `json:"ambient_temperature_c,omitempty"`
	AmbientTemperatureF       int       `json:"ambient_temperature_f,omitempty"`
	AwayTemperatureHighC      float64   `json:"away_temperature_high_c,omitempty"`
	AwayTemperatureHighF      int       `json:"away_temperature_high_f,omitempty"`
	AwayTemperatureLowC       float64   `json:"away_temperature_low_c,omitempty"`
	AwayTemperatureLowF       int       `json:"away_temperature_low_f,omitempty"`
	EcoTemperatureHighC       float64   `json:"eco_temperature_high_c,omitempty"`
	EcoTemperatureHighF       int       `json:"eco_temperature_high_f,omitempty"`
	EcoTemperatureLowC        float64   `json:"eco_temperature_low_c,omitempty"`
	EcoTemperatureLowF        int       `json:"eco_temperature_low_f,omitempty"`
	IsLocked                  bool      `json:"is_locked,omitempty"`
	LockedTempMinC            float64   `json:"locked_temp_min_c,omitempty"`
	LockedTempMinF            int       `json:"locked_temp_min_f,omitempty"`
	LockedTempMaxC            float64   `json:"locked_temp_max_cv,omitempty"`
	LockedTempMaxF            int       `json:"locked_temp_max_f,omitempty"`
	SunlightCorrectionActive  bool      `json:"sunlight_correction_active,omitempty"`
	SunlightCorrectionEnabled bool      `json:"sunlight_correction_enabled,omitempty"`
	StructureID               string    `json:"structure_id,omitempty"`
	FanTimerActive            bool      `json:"fan_timer_active,omitempty"`
	FanTimerTimeout           time.Time `json:"fan_timer_timeout,omitempty"`
	FanTimerDuration          int       `json:"fan_timer_duration,omitempty"`
	PreviousHvacMode          string    `json:"previous_hvac_mode,omitempty"`
	HvacMode                  string    `json:"hvac_mode,omitempty"`
	TimeToTarget              string    `json:"time_to_target,omitempty"`
	TimeToTargetTraining      string    `json:"time_to_target_training,omitempty"`
	WhereName                 string    `json:"where_name,omitempty"`
	Label                     string    `json:"label,omitempty"`
	NameLong                  string    `json:"name_long,omitempty"`
	IsOnline                  bool      `json:"is_online,omitempty"`
	LastConnection            time.Time `json:"last_connection,omitempty"`
	HvacState                 string    `json:"hvac_state,omitempty"`
}
