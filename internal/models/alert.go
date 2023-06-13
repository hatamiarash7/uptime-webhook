package models

import "time"

// Alert is a struct for alert model
type Alert struct {
	Event string    `json:"event"`
	Data  AlertData `json:"data"`
}

// AlertData is a struct for alert data model
type AlertData struct {
	Service          ServiceInfo          `json:"service"`
	Account          AccountInfo          `json:"account"`
	Integration      IntegrationInfo      `json:"integration"`
	Date             time.Time            `json:"date"`
	Alert            AlertInfo            `json:"alert"`
	GlobalAlertState GlobalAlertStateInfo `json:"global_alert_state"`
	Device           DeviceInfo           `json:"device"`
	Locations        []string             `json:"locations"`
	Links            LinksInfo            `json:"links"`
}

// ServiceInfo is a struct for service info model in Alert's data
type ServiceInfo struct {
	ID                           int      `json:"id"`
	Name                         string   `json:"name"`
	DeviceID                     int      `json:"device_id"`
	MonitoringServiceType        string   `json:"monitoring_service_type"`
	IsPaused                     bool     `json:"is_paused"`
	MspAddress                   string   `json:"msp_address"`
	MspVersion                   int      `json:"msp_version"`
	MspInterval                  int      `json:"msp_interval"`
	MspSensitivity               int      `json:"msp_sensitivity"`
	MspNumRetries                int      `json:"msp_num_retries"`
	MspURLScheme                 string   `json:"msp_url_scheme"`
	MspURLPath                   string   `json:"msp_url_path"`
	MspPort                      any      `json:"msp_port"`
	MspUsername                  string   `json:"msp_username"`
	MspProxy                     string   `json:"msp_proxy"`
	MspDNSServer                 string   `json:"msp_dns_server"`
	MspDNSRecordType             string   `json:"msp_dns_record_type"`
	MspStatusCode                string   `json:"msp_status_code"`
	MspSendString                string   `json:"msp_send_string"`
	MspExpectString              string   `json:"msp_expect_string"`
	MspExpectStringType          string   `json:"msp_expect_string_type"`
	MspEncryption                string   `json:"msp_encryption"`
	MspThreshold                 int      `json:"msp_threshold"`
	MspNotes                     string   `json:"msp_notes"`
	MspIncludeInGlobalMetrics    bool     `json:"msp_include_in_global_metrics"`
	MspUseIPVersion              string   `json:"msp_use_ip_version"`
	MspUptimeSLA                 string   `json:"msp_uptime_sla"`
	MspResponseTimeSLA           string   `json:"msp_response_time_sla"`
	MonitoringServiceTypeDisplay string   `json:"monitoring_service_type_display"`
	DisplayName                  string   `json:"display_name"`
	ShortName                    string   `json:"short_name"`
	Tags                         []string `json:"tags"`
}

// AccountInfo is a struct for account info model in Alert's data
type AccountInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Brand    string `json:"brand"`
	Timezone string `json:"timezone"`
	SiteURL  string `json:"site_url"`
}

// IntegrationInfo is a struct for integration info model in Alert's data
type IntegrationInfo struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Module            string `json:"module"`
	ModuleVerboseName string `json:"module_verbose_name"`
	IsEnabled         bool   `json:"is_enabled"`
	IsErrored         bool   `json:"is_errored"`
	IsTestSupported   bool   `json:"is_test_supported"`
	PostbackURL       string `json:"postback_url"`
	Headers           string `json:"headers"`
	UseLegacyPayload  bool   `json:"use_legacy_payload"`
}

// AlertInfo is a struct for alert info model in Alert's data
type AlertInfo struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	State       string    `json:"state"`
	Output      string    `json:"output"`
	ShortOutput string    `json:"short_output"`
	IsUp        bool      `json:"is_up"`
}

// GlobalAlertStateInfo is a struct for global alert state info model in Alert's data
type GlobalAlertStateInfo struct {
	ID               int       `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	NumLocationsDown int       `json:"num_locations_down"`
	StateIsUp        bool      `json:"state_is_up"`
	StateHasChanged  bool      `json:"state_has_changed"`
	Ignored          bool      `json:"ignored"`
}

// DeviceInfo is a struct for device info model in Alert's data
type DeviceInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	IsPaused    bool   `json:"is_paused"`
	DisplayName string `json:"display_name"`
}

// LinksInfo is a struct for links info model in Alert's data
type LinksInfo struct {
	AlertDetails     string `json:"alert_details"`
	RealTimeAnalysis string `json:"real_time_analysis"`
}
