package emojies

import "main/internal/models"

type EmojiUse interface {
	CreateEmoji(models.Emoji) error
}
