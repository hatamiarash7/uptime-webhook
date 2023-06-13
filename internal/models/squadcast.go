package models

// SquadcastIncident is a struct for Squadcast incident
type SquadcastIncident struct {
	Message     string                  `json:"message,omitempty"`
	Description string                  `json:"description,omitempty"`
	Tags        map[string]SquadcastTag `json:"tags,omitempty"`
	Status      string                  `json:"status"`
	EventID     string                  `json:"event_id"`
	Locations   string                  `json:"locations"`
	Retries     string                  `json:"retries"`
	Type        string                  `json:"type"`
}

// SquadcastTag is a struct for Squadcast tag
type SquadcastTag struct {
	Color string `json:"color"`
	Value string `json:"value"`
}
