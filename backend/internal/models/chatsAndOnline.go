package models

type ChatAndOnline struct {
	Online []OnlineUsers `json:"online"`
	Chats  []Chat        `json:"chats"`
}
