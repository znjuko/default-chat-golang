package models

type Chat struct {
	ChatId              *string `json:"chatId"`
	ChatName            *string `json:"chatName"`
	ChatLastMessage     *string `json:"chatMsg"`
	ChatLastAuthorLogin *string `json:"chatAuthor"`
}
