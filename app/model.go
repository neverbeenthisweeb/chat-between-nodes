package app

import "time"

type Chat struct {
	ID        string    `json:"id,omitempty"`
	NodeID    string    `json:"node_id,omitempty"` // ID to differentiate browser tabs
	From      string    `json:"from,omitempty"`    // Sender's user name
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
