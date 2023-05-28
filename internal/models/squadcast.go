package models

type SquadcastIncident struct {
	Message     string                  `json:"message,omitempty"`
	Description string                  `json:"description,omitempty"`
	Tags        map[string]SquadcastTag `json:"tags,omitempty"`
	Status      string                  `json:"status"`
	EventID     string                  `json:"event_id"`
}

type SquadcastTag struct {
	Color string `json:"color"`
	Value string `json:"value"`
}
