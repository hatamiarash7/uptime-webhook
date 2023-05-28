package alert

import (
	"encoding/json"
	"time"
)

type CreateAlertRequest struct {
	Event string `json:"event" binding:"required"`
	Data  struct {
		Service struct {
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
		} `json:"service"`
		Account struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Brand    string `json:"brand"`
			Timezone string `json:"timezone"`
			SiteURL  string `json:"site_url"`
		} `json:"account"`
		Integration struct {
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
		} `json:"integration"`
		Date  time.Time `json:"date"`
		Alert struct {
			ID          int       `json:"id"`
			CreatedAt   time.Time `json:"created_at"`
			State       string    `json:"state"`
			Output      string    `json:"output"`
			ShortOutput string    `json:"short_output"`
			IsUp        bool      `json:"is_up"`
		} `json:"alert"`
		GlobalAlertState struct {
			ID               int       `json:"id"`
			CreatedAt        time.Time `json:"created_at"`
			NumLocationsDown int       `json:"num_locations_down"`
			StateIsUp        bool      `json:"state_is_up"`
			StateHasChanged  bool      `json:"state_has_changed"`
			Ignored          bool      `json:"ignored"`
		} `json:"global_alert_state"`
		Device struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Address     string `json:"address"`
			IsPaused    bool   `json:"is_paused"`
			DisplayName string `json:"display_name"`
		} `json:"device"`
		Locations []string `json:"locations"`
		Links     struct {
			AlertDetails     string `json:"alert_details"`
			RealTimeAnalysis string `json:"real_time_analysis"`
		} `json:"links"`
	} `json:"data,omitempty" binding:"required"`
}

func (h *CreateAlertRequest) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}
