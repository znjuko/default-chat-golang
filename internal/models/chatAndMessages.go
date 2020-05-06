package models

type ChatAndMsgs struct {
	ChatInfo Chat `json:"chatInfo,omitempty"`
	Messages []Message `json:"messages"`
}