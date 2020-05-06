package models

type Chat struct {
	ChatId              *string `json:"chatId,omitempty"`
	ChatName            *string `json:"chatName,omitempty"`
	ChatLastMessage     *string `json:"chatMsg,omitempty"`
	ChatLastAuthorLogin *string `json:"chatAuthor,omitempty"`
}
