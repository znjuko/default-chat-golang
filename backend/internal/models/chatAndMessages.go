package models

type ChatAndMsgs struct {
	ChatInfo   Chat      `json:"chatInfo"`
	Messages   []Message `json:"messages"`
	AllEmojies []Emoji   `json:"allEmojies"`
}
