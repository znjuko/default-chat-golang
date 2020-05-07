package models

type Chat struct {
	ChatId              int `json:"chatId"`
	ChatName            *string `json:"chatName"`
	ChatLastMessage     *string `json:"chatMsg"`
	ChatLastAuthorLogin *string `json:"chatAuthor"`
}
