package models

type NewChatUsers struct {
	ChatName  *string `json:"chatName"`
	ChatUsers []int   `json:"users"`
}
