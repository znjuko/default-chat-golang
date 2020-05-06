package models

type Session struct {
	Login *string `json:"login"`
	Password *string `json:"password"`
}