package models

type Chat struct {
	ChatName            *string `json:"chatName,omitempty"`
	ChatLastMessage     *string `json:"chatMsg,omitempty"`
	ChatLastAuthorLogin *string `json:"chatAuthor,omitempty"`
}
