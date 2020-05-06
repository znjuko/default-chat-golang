package models

type Message struct {
	AuthorLogin *string `json:"author,omitempty"`
	Text        *string `json:"txt,omitempty"`
	ChatName    *string `json:"chatName,omitempty"`
	ChatId      *string `json:"chatId,omitempty"`
}
