package ws

import "encoding/json"

type Event struct {
	Type    string     `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client, m *Manager) error

const (
	EventSendMessage = "send_message"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
	To 		string `json:"to"` 
}