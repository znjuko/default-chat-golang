package usecase

import (
	"main/internal/message"
	stocken "main/internal/socket_token"
	cr "main/internal/tools/connection_reciever"
	"main/internal/tools/eventer"
	"net"
	"main/internal/socket"
)

type SocketUseCaseRealisation struct {
	messageDB message.MessageRepository
	tokenDB   stocken.TokenRepository
	onlineDB socket.OnlineRepo
}

func NewSocketUseCaseRealisation(messageDB message.MessageRepository, tokenDB stocken.TokenRepository , online socket.OnlineRepo) SocketUseCaseRealisation {
	return SocketUseCaseRealisation{messageDB: messageDB, tokenDB: tokenDB , onlineDB: online}
}

func (SU SocketUseCaseRealisation) CheckToken(tokenValue string) (int, error) {
	return SU.tokenDB.GetUserIdByToken(tokenValue)
}

func (SU SocketUseCaseRealisation) AddToConnectionPool(conn net.Conn, userId int) error {

	eventer := eventer.NewEventer(userId, SU.messageDB, SU.onlineDB)

	reciever, err := cr.NewConnReciever(conn, eventer)
	if err == nil {
		reciever.StartRecieving()
		SU.onlineDB.AddOnline(userId)
	}
	return err
}
