package models

type ChatAndOnline struct {
	Online []OnlineUsers `json:"online,omitempty"`
	Chats  []Chat        `json:"chats,omitempty"`
}
