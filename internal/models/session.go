package models

type Session struct {
	Login *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
}