package models

type JSONEvent struct {
	Event   string        `json:"event"`
	Message Message       `json:"message"`
	Online  []OnlineUsers `json:"online"`
}
