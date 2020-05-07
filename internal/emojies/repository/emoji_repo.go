package repository

import (
	"database/sql"
	"main/internal/models"
)

type EmojiRepoRealisation struct {
	database *sql.DB
}

func NewEmojiUseRealisation(db *sql.DB) EmojiRepoRealisation {
	return EmojiRepoRealisation{database: db}
}

func (Emoji EmojiRepoRealisation) CreateEmoji(emoji models.Emoji) error {
	_, err := Emoji.database.Exec("INSERT INTO emoji (main_word,slug) VALUES($1,$2)", *emoji.Phrase, *emoji.Url)

	return err
}

func (Emoji EmojiRepoRealisation) GetAllEmojies() ([]models.Emoji, error) {
	emRows, err := Emoji.database.Query("SELECT main_word , slug FROM emoji")

	defer func() {
		if emRows != nil {
			emRows.Close()
		}
	}()

	emjs := make([]models.Emoji, 0)

	for emRows.Next() {
		emj := new(models.Emoji)
		err = emRows.Scan(&emj.Phrase, &emj.Url)

		if err != nil {
			return nil, err
		}

		emjs = append(emjs, *emj)

	}

	return emjs, nil
}
