package repository

import (
	"database/sql"
	"main/internal/models"
)

type EmojiRepoRealisation struct{
	database *sql.DB
}

func NewEmojiUseRealisation(db *sql.DB) EmojiRepoRealisation {
	return EmojiRepoRealisation{database: db}
}

func (Emoji EmojiRepoRealisation) CreateEmoji(emoji models.Emoji) error {
	_ , err := Emoji.database.Exec("INSERT INTO emoji (main_word,slug) VALUES($1,$2)", *emoji.Phrase, *emoji.Url)

	return err
}
