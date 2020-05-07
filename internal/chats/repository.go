package chats

import "main/internal/models"

type ChatRepo interface {
	GetOnline(int) ([]models.OnlineUsers, error)
	GetChats(int) ([]models.Chat, error)
	CreateChat(string, []int) error
	GetChat(int) (models.Chat, []models.Message,[]models.Emoji ,error)
}
