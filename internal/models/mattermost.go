package models

// MattermostMessage is a struct for Mattermost message
type MattermostMessage struct {
	Text     string `json:"text,omitempty"`
	Username string `json:"username,omitempty"`
	Icon     string `json:"icon_url,omitempty"`
}
