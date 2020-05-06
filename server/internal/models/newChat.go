package models

type NewChatUsers struct {
	ChatName  *string `json:"chatName,omitempty"`
	ChatUsers []int   `json:"users,omitempty"`
}
