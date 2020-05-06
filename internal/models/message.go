package models

type Message struct {
	AuthorLogin *string `json:"author"`
	Text        *string `json:"txt"`
	ChatName    *string `json:"chatName"`
	ChatId      *string `json:"chatId"`
	Emojies     []Emoji `json:"emojies"`
}
