package socket

import "main/internal/models"

type OnlineRepo interface {
	AddOnline(int) error
	DiscardOnline(int) error
	GetOnline(int) ([]models.OnlineUsers, error)
}
