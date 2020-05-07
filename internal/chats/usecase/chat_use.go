package usecase

import (
	"main/internal/chats"
	"main/internal/emojies"
	"main/internal/models"
)

type ChatUseCaseRealisation struct {
	emojiDB emojies.EmojiRepo
	chatDB  chats.ChatRepo
}

func NewChatUseCaseRealisation(chat chats.ChatRepo, emoji emojies.EmojiRepo) ChatUseCaseRealisation {
	return ChatUseCaseRealisation{chatDB: chat, emojiDB: emoji}
}

func (Chat ChatUseCaseRealisation) CreateChat(chat models.NewChatUsers, userId int) error {

	chat.ChatUsers = append(chat.ChatUsers, userId)

	return Chat.chatDB.CreateChat(*chat.ChatName, chat.ChatUsers)

}

func (Chat ChatUseCaseRealisation) GetChatsAndOnlineUsers(userId int) (models.ChatAndOnline, error) {

	var err error
	chAndOnl := new(models.ChatAndOnline)

	chAndOnl.Online, err = Chat.chatDB.GetOnline(userId)

	if err != nil {
		return *chAndOnl, err
	}

	chAndOnl.Chats, err = Chat.chatDB.GetChats(userId)

	return *chAndOnl, err

}

func (Chat ChatUseCaseRealisation) GetChat(chatId int) (models.ChatAndMsgs, error) {
	chAndMsg := new(models.ChatAndMsgs)
	var err error
	chAndMsg.ChatInfo, chAndMsg.Messages, err = Chat.chatDB.GetChat(chatId)

	if err != nil {
		return models.ChatAndMsgs{}, err
	}

	chAndMsg.AllEmojies, err = Chat.emojiDB.GetAllEmojies()

	return *chAndMsg, err
}
