package usecase

import (
	"main/internal/chats"
	"main/internal/models"
)

type ChatUseCaseRealisation struct {
	chatDB chats.ChatRepo
}

func NewChatUseCaseRealisation(chat chats.ChatRepo) ChatUseCaseRealisation {
	return ChatUseCaseRealisation{chatDB: chat}
}

func (Chat ChatUseCaseRealisation) CreateChat(chat models.NewChatUsers, userId int) error {

	chat.ChatUsers = append(chat.ChatUsers, userId)

	return Chat.chatDB.CreateChat(*chat.ChatName, chat.ChatUsers)

}

func (Chat ChatUseCaseRealisation) GetChatsAndOnlineUsers(userId int) (models.ChatAndOnline, error) {

	var err error
	chAndOnl := new(models.ChatAndOnline)

	chAndOnl.Online, err = Chat.chatDB.GetOnline()

	if err != nil {
		return *chAndOnl, err
	}

	chAndOnl.Chats, err = Chat.chatDB.GetChats(userId)

	return *chAndOnl, err

}

func (Chat ChatUseCaseRealisation) GetChat(chatId int) (models.ChatAndMsgs , error) {
	chAndMsg := new(models.ChatAndMsgs)
	var err error
	chAndMsg.ChatInfo , chAndMsg.Messages , err = Chat.chatDB.GetChat(chatId)

	return *chAndMsg , err
}
