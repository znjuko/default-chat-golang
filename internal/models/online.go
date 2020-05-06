package models

type OnlineUsers struct {
	Login  *string `json:"login,omitempty"`
	UserId *int    `json:"id,omitempty"`
}
