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
