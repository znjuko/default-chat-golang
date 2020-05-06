package models

type Emoji struct {
	Url *string `json:"url,omitempty"`
	Phrase *string `json:"phrase,omitempty"`
}
