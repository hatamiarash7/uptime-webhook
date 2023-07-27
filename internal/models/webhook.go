package models

// CustomMessage is a struct for Telegram message
type CustomMessage struct {
	Status      string `json:"status,omitempty"`
	ShortName   string `json:"short_name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Date        string `json:"date,omitempty"`
	Address     string `json:"address,omitempty"`
	ShortOutput string `json:"short_output,omitempty"`
}
