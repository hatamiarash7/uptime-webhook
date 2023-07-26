package monitoring

import (
	log "github.com/sirupsen/logrus"
)

// MockMonitor is a mock monitor
type MockMonitor struct {
}

// NewMockMonitor creates a new mock monitor
func NewMockMonitor() Monitor {
	return MockMonitor{}
}

// Record records the events
func (i MockMonitor) Record(events []Event) {
	for _, event := range events {
		switch event.GetID() {
		case IncTotalAlert:
			log.Info("====IncTotalAlert=====")
		case IncTelegramSendSuccess:
			log.Info("====IncTelegramSendSuccess=====")
		case IncTelegramSendFailure:
			log.Info("====IncTelegramSendFailure=====")
		case IncSquadcastSendSuccess:
			log.Info("====IncSquadcastSendSuccess=====")
		case IncSquadcastSendFailure:
			log.Info("====IncSquadcastSendFailure=====")
		case IncSlackSendSuccess:
			log.Info("====IncSlackSendSuccess=====")
		case IncSlackSendFailure:
			log.Info("====IncSlackSendFailure=====")
		case SetActiveJobsInAlertPool:
			log.Info("====SetActiveJobsInAlertPool=====")
		case SetAlertPoolCapacity:
			log.Info("====SetAlertPoolCapacity=====")
		default:
			log.Errorf("[MONITORING] Invalid event id [%d]", event.GetID())
		}
	}
}
