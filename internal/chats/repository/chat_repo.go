package repository

import (
	"database/sql"
	"main/internal/models"
)

type ChatRepoRealisation struct {
	database *sql.DB
}

func NewChatRepoRealistaion(db *sql.DB) ChatRepoRealisation {
	return ChatRepoRealisation{database: db}
}

func (Chat ChatRepoRealisation) GetOnline(userId int) ([]models.OnlineUsers, error) {

	row, err := Chat.database.Query("SELECT U.u_id , U.login FROM users U INNER JOIN online O ON (O.u_id=U.u_id) WHERE U.u_id != $1", userId)

	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	users := make([]models.OnlineUsers, 0)

	for row.Next() {
		user := new(models.OnlineUsers)

		err = row.Scan(&user.UserId, &user.Login)

		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return users, nil

}

func (Chat ChatRepoRealisation) GetChat(chatId int) (models.Chat, []models.Message, []models.Emoji, error) {

	chat := new(models.Chat)
	row := Chat.database.QueryRow("SELECT name FROM chats WHERE ch_id = $1", chatId)
	err := row.Scan(&chat.ChatName)

	if err != nil {
		return *chat, nil, nil, err
	}

	rows, err := Chat.database.Query("SELECT M.msg_id , M.text , U.login FROM messages M INNER JOIN users U ON(M.u_id=U.u_id) WHERE M.ch_id = $1 ORDER BY M.msg_id DESC", chatId)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return *chat, nil, nil, err
	}

	msgs := make([]models.Message, 0)

	for rows.Next() {
		msg := new(models.Message)
		var msgid *int
		err = rows.Scan(&msgid, &msg.Text, &msg.AuthorLogin)

		if err != nil {
			return *chat, nil, nil, err
		}

		msgs = append(msgs, *msg)
	}

	emRows, err := Chat.database.Query("SELECT main_word , slug FROM emoji")

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
			return *chat, nil, nil, err
		}

		emjs = append(emjs, *emj)

	}

	return *chat, msgs, emjs, nil

}

func (Chat ChatRepoRealisation) GetChats(userId int) ([]models.Chat, error) {

	row, err := Chat.database.Query("SELECT C.ch_id,C.name , C.last_msg_id , C.last_msg_log , C.last_msg_txt FROM chats C INNER JOIN chat_user CU ON (C.ch_id=CU.ch_id) WHERE CU.u_id = $1 ORDER BY C.last_msg_id DESC", userId)
	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	chats := make([]models.Chat, 0)

	for row.Next() {
		chat := new(models.Chat)
		msgId := 0

		err = row.Scan(&chat.ChatId, &chat.ChatName, &msgId, &chat.ChatLastAuthorLogin, &chat.ChatLastMessage)

		if err != nil {
			return nil, err
		}

		chats = append(chats, *chat)

	}

	return chats, err
}

func (Chat ChatRepoRealisation) CreateChat(chatName string, chatUsers []int) error {

	chatId := 0
	row := Chat.database.QueryRow("INSERT INTO chats (name) VALUES($1) RETURNING ch_id", chatName)

	err := row.Scan(&chatId)

	if err != nil {
		return err
	}

	for _, value := range chatUsers {
		_, err = Chat.database.Exec("INSERT INTO chat_user (ch_id,u_id) VALUES($1,$2)", chatId, value)

		if err != nil {
			return err
		}
	}

	return nil
}
