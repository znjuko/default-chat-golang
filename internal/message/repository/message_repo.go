package repository

import (
	"database/sql"
	"fmt"
	"main/internal/models"
	"strings"
)

type MessageRepositoryRealisation struct {
	messageDB *sql.DB
}

func NewMessageRepositoryRealisation(addr, pass string, db *sql.DB) MessageRepositoryRealisation {

	return MessageRepositoryRealisation{messageDB: db}
}

func MessageParser(messageTxt string) []string {

	emojies := make([]string, 0)
	letters := strings.Fields(messageTxt)

	for _, value := range letters {
		if value[:1] == ":" && value[len(value)-1:] == ":" {
			emojies = append(emojies, value)
		}
	}

	return emojies
}

func (MR MessageRepositoryRealisation) AddNewMessage(author int, message models.Message) error {

	msgId := 0
	row := MR.messageDB.QueryRow("INSERT INTO messages (u_id,ch_id,text) VALUES($1,$2,$3) RETURNING msg_id", author, message.ChatId, *message.Text)
	err := row.Scan(&msgId)
	if err != nil {
		return err
	}

	var usrLogin *string
	row = MR.messageDB.QueryRow("SELECT login FROM users WHERE u_id = $1", author)
	err = row.Scan(&usrLogin)
	if err != nil {
		return err
	}

	MR.messageDB.Exec("UPDATE chats SET last_msg_id = $1 , last_msg_log = $2 , last_msg_txt = $3 WHERE ch_id = $4", msgId, *usrLogin, *message.Text, message.ChatId)

	recRows, err := MR.messageDB.Query("SELECT u_id FROM chat_user WHERE ch_id = $1", message.ChatId)
	defer func() {
		if recRows != nil {
			recRows.Close()
		}
	}()

	if err != nil {
		return err
	}

	for recRows.Next() {
		reciever := 0

		err = recRows.Scan(&reciever)
		_, err = MR.messageDB.Exec("INSERT INTO newmessages (msg_id,u_id) VALUES($1,$2)", msgId, reciever)

	}

	return nil
}

func (MR MessageRepositoryRealisation) ReceiveNewMessages(userId int) ([]models.Message, error) {

	msgsArray := make([]models.Message, 0)

	msgsRow, err := MR.messageDB.Query(
		"SELECT M.msg_id,M.text,U.login,C.ch_id,C.name FROM messages M INNER JOIN newmessages NM ON(NM.msg_id=M.msg_id) "+
			" INNER JOIN users U ON(M.u_id=U.u_id) INNER JOIN chats C ON(C.ch_id=M.ch_id) "+
			" WHERE NM.u_id = $1",
		userId)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for msgsRow.Next() {

		msgId := 0
		msg := new(models.Message)

		err = msgsRow.Scan(&msgId, &msg.Text, &msg.AuthorLogin, &msg.ChatId, &msg.ChatName)

		if err != nil {
			return nil, err
		}

		MR.messageDB.Exec("DELETE FROM newmessages WHERE msg_id = $1 AND u_id = $2", msgId, userId)

		emojies := MessageParser(*msg.Text)

		if len(emojies) != 0 {
			msg.Emojies = make([]models.Emoji, 0)
			for iter , value := range emojies {

				row := MR.messageDB.QueryRow("SELECT slug FROM emoji WHERE main_word = $1", value)
				var emojiSlug *string

				err := row.Scan(&emojiSlug)

				if err == nil {
					msg.Emojies = append(msg.Emojies, models.Emoji{
						Url:    emojiSlug,
						Phrase: &emojies[iter],
					})
				}


			}
		}

		msgsArray = append(msgsArray, *msg)

	}

	return msgsArray, nil
}
