package action

import (
	requests "github.com/arvancloud/uptime-webhook/internal/http/requests/alert"
	"github.com/arvancloud/uptime-webhook/internal/models"
)

func transformRequestToValue(c requests.CreateAlertRequest) models.Alert {
	h := models.Alert{
		Event: c.Event,
		Data: models.AlertData{
			Service: models.ServiceInfo{
				ID:                           c.Data.Service.ID,
				Name:                         c.Data.Service.Name,
				DeviceID:                     c.Data.Service.DeviceID,
				MonitoringServiceType:        c.Data.Service.MonitoringServiceType,
				IsPaused:                     c.Data.Service.IsPaused,
				MspAddress:                   c.Data.Service.MspAddress,
				MspVersion:                   c.Data.Service.MspVersion,
				MspInterval:                  c.Data.Service.MspInterval,
				MspSensitivity:               c.Data.Service.MspSensitivity,
				MspNumRetries:                c.Data.Service.MspNumRetries,
				MspURLScheme:                 c.Data.Service.MspURLScheme,
				MspURLPath:                   c.Data.Service.MspURLPath,
				MspPort:                      c.Data.Service.MspPort,
				MspUsername:                  c.Data.Service.MspUsername,
				MspProxy:                     c.Data.Service.MspProxy,
				MspDNSServer:                 c.Data.Service.MspDNSServer,
				MspDNSRecordType:             c.Data.Service.MspDNSRecordType,
				MspStatusCode:                c.Data.Service.MspStatusCode,
				MspSendString:                c.Data.Service.MspSendString,
				MspExpectString:              c.Data.Service.MspExpectString,
				MspExpectStringType:          c.Data.Service.MspExpectStringType,
				MspEncryption:                c.Data.Service.MspEncryption,
				MspThreshold:                 c.Data.Service.MspThreshold,
				MspNotes:                     c.Data.Service.MspNotes,
				MspIncludeInGlobalMetrics:    c.Data.Service.MspIncludeInGlobalMetrics,
				MspUseIPVersion:              c.Data.Service.MspUseIPVersion,
				MspUptimeSLA:                 c.Data.Service.MspUptimeSLA,
				MspResponseTimeSLA:           c.Data.Service.MspResponseTimeSLA,
				MonitoringServiceTypeDisplay: c.Data.Service.MonitoringServiceTypeDisplay,
				DisplayName:                  c.Data.Service.DisplayName,
				ShortName:                    c.Data.Service.ShortName,
				Tags:                         c.Data.Service.Tags,
			},
			Account: models.AccountInfo{
				ID:       c.Data.Account.ID,
				Name:     c.Data.Account.Name,
				Brand:    c.Data.Account.Brand,
				Timezone: c.Data.Account.Timezone,
				SiteURL:  c.Data.Account.SiteURL,
			},
			Integration: models.IntegrationInfo{
				ID:                c.Data.Integration.ID,
				Name:              c.Data.Integration.Name,
				Module:            c.Data.Integration.Module,
				ModuleVerboseName: c.Data.Integration.ModuleVerboseName,
				IsEnabled:         c.Data.Integration.IsEnabled,
				IsErrored:         c.Data.Integration.IsErrored,
				IsTestSupported:   c.Data.Integration.IsTestSupported,
				PostbackURL:       c.Data.Integration.PostbackURL,
				Headers:           c.Data.Integration.Headers,
				UseLegacyPayload:  c.Data.Integration.UseLegacyPayload,
			},
			Date: c.Data.Date,
			Alert: models.AlertInfo{
				ID:          c.Data.Alert.ID,
				CreatedAt:   c.Data.Alert.CreatedAt,
				State:       c.Data.Alert.State,
				Output:      c.Data.Alert.Output,
				ShortOutput: c.Data.Alert.ShortOutput,
				IsUp:        c.Data.Alert.IsUp,
			},
			GlobalAlertState: models.GlobalAlertStateInfo{
				ID:               c.Data.GlobalAlertState.ID,
				CreatedAt:        c.Data.GlobalAlertState.CreatedAt,
				NumLocationsDown: c.Data.GlobalAlertState.NumLocationsDown,
				StateIsUp:        c.Data.GlobalAlertState.StateIsUp,
				StateHasChanged:  c.Data.GlobalAlertState.StateHasChanged,
				Ignored:          c.Data.GlobalAlertState.Ignored,
			},
			Device: models.DeviceInfo{
				ID:          c.Data.Device.ID,
				Name:        c.Data.Device.Name,
				Address:     c.Data.Device.Address,
				IsPaused:    c.Data.Device.IsPaused,
				DisplayName: c.Data.Device.DisplayName,
			},
			Locations: c.Data.Locations,
			Links: models.LinksInfo{
				AlertDetails:     c.Data.Links.AlertDetails,
				RealTimeAnalysis: c.Data.Links.RealTimeAnalysis,
			},
		},
	}
	return h
}
