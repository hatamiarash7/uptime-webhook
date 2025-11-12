package monitoring

// EventName is the name of event
type EventName uint

const (
	// IncTotalAlert will show the total incoming alerts
	IncTotalAlert EventName = iota
	// IncTelegramSendSuccess will show the total success telegram sends
	IncTelegramSendSuccess
	// IncTelegramSendFailure will show the total failure telegram sends
	IncTelegramSendFailure
	// IncSquadcastSendSuccess will show the total success squadcast sends
	IncSquadcastSendSuccess
	// IncSquadcastSendFailure will show the total failure squadcast sends
	IncSquadcastSendFailure
	// IncMattermostSendSuccess will show the total success squadcast sends
	IncMattermostSendSuccess
	// IncMattermostSendFailure will show the total failure squadcast sends
	IncMattermostSendFailure
	// IncSlackSendSuccess will show the total success squadcast sends
	IncSlackSendSuccess
	// IncSlackSendFailure will show the total failure squadcast sends
	IncSlackSendFailure
	// IncCustomSendSuccess will show the total success squadcast sends
	IncCustomSendSuccess
	// IncCustomSendFailure will show the total failure squadcast sends
	IncCustomSendFailure
	// SetActiveJobsInAlertPool will show the total active jobs in alert pool
	SetActiveJobsInAlertPool
	// SetAlertPoolCapacity will show the total capacity of alert pool
	SetAlertPoolCapacity
	// SetCheckStatus will set the status of a check
	SetCheckStatus
)

// Event is the event structure
type Event struct {
	id     EventName
	params []interface{}
}

// GetID returns the event's id
func (e Event) GetID() EventName {
	return e.id
}

// GetParam returns the event's parameter
func (e Event) GetParam(i int) interface{} {
	if i >= len(e.params) {
		return nil
	}
	return e.params[i]
}

// NewEvent creates a new event
func NewEvent(name EventName, params ...interface{}) Event {
	return Event{id: name, params: params}
}
