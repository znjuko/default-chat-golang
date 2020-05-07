package eventer

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"main/internal/message"
	"main/internal/models"
	"main/internal/socket"
	"net"
)

type Eventer struct {
	userId          int
	messageDB       message.MessageRepository
	onlineDiscarded socket.OnlineRepo
}

func NewEventer(user int, messages message.MessageRepository, online socket.OnlineRepo) Eventer {
	return Eventer{userId: user, messageDB: messages, onlineDiscarded: online}
}

func (EV Eventer) WriteNewMessage(conn net.Conn) {

	fmt.Println("wrote new message")

	req := wsutil.NewReader(conn, ws.StateServerSide)
	decoder := json.NewDecoder(req)
	hdr, err := req.NextFrame()
	resp := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)

	if err != nil || hdr.OpCode == ws.OpClose {
		fmt.Println("connection closed")
		conn.Close()
		return
	}

	defer func() {
		resp.Flush()
	}()

	msg := new(models.Message)

	if err = decoder.Decode(&msg); err != nil {
		fmt.Println("DECODE JSON ERROR : ", err)
		return
	}

	if err := EV.messageDB.AddNewMessage(EV.userId, *msg); err != nil {
		fmt.Println("ADD NEW MESSAGES ERROR : ", err)
		return
	}

}

func (EV Eventer) GetNewMessages(conn net.Conn) {
	resp := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(resp)
	answer := models.JSONEvent{}

	defer func() {
		resp.Flush()
	}()
	messages, err := EV.messageDB.ReceiveNewMessages(EV.userId)

	if err != nil {
		answer.Event = "can't get new messages"
		encoder.Encode(&answer)
		fmt.Println("GET NEW MESSAGES ERROR : ", err)
		return
	}

	answer.Event = "new message"

	for iter, _ := range messages {
		answer.Message = messages[iter]

		for i , _ := range answer.Message.Emojies {
			fmt.Println("sended emojies : ")
			fmt.Println(*answer.Message.Emojies[i].Phrase)
		}

		encoder.Encode(&answer)
		resp.Flush()
	}

}

func (EV Eventer) GetOnline(conn net.Conn) {
	resp := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(resp)
	answer := models.JSONEvent{}

	onlines, err := EV.onlineDiscarded.GetOnline(EV.userId)


	if err != nil {
		answer.Event = "can't get online"
		encoder.Encode(&answer)
		fmt.Println("GET ONLINE ERROR : ", err)
		return
	}

	answer.Event = "online"
	answer.Online = onlines

	encoder.Encode(&answer)
	resp.Flush()
}

func (EV Eventer) DiscardOnline() {
	EV.onlineDiscarded.DiscardOnline(EV.userId)
}
