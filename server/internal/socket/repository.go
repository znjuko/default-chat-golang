package socket

type OnlineRepo interface {
	AddOnline(int) error
	DiscardOnline(int) error
}
