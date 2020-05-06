package usecase

import (
	"main/internal/emojies"
	"main/internal/models"
)

type EmojiUseRealisation struct{
	emojiDB emojies.EmojiRepo
}

func NewEmojiUseRealisation(emoji emojies.EmojiRepo) EmojiUseRealisation {
	return EmojiUseRealisation{emojiDB: emoji}
}

func (Emoji EmojiUseRealisation) CreateEmoji(emoji models.Emoji) error {
	return Emoji.emojiDB.CreateEmoji(emoji)
}
