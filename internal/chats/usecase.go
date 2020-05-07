package chats

import "main/internal/models"

type ChatUseCase interface {
	CreateChat(models.NewChatUsers, int) error
	GetChatsAndOnlineUsers(int) (models.ChatAndOnline, error)
	GetChat(int) (models.ChatAndMsgs, error)
}
