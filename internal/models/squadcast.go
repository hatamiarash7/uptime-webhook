package models

type SquadcastTrigger struct {
	Message     string                  `json:"message"`
	Description string                  `json:"description"`
	Tags        map[string]SquadcastTag `json:"tags"`
	Status      string                  `json:"status"`
	EventID     string                  `json:"event_id"`
}

type SquadcastTag struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

type SquadcastResolve struct {
	Status  string `json:"status"`
	EventID string `json:"event_id"`
}
