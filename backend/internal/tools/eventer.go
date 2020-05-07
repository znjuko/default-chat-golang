package tools

import "net"

type EventerInterface interface {
	GetNewMessages(net.Conn)
	WriteNewMessage(net.Conn)
	DiscardOnline()
	GetOnline(conn net.Conn)
}
