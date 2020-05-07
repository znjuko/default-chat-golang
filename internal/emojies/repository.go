package emojies

import "main/internal/models"

type EmojiRepo interface {
	CreateEmoji(models.Emoji) error
	GetAllEmojies() ([]models.Emoji, error)
}
