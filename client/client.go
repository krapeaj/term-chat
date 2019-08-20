package main

type Client interface {
	Login(userId, password string) error
	Logout() error
	Create() error
	Delete(chatId string) error
	Enter(chatId string) error
	Leave() error
	SendMessage(message string) error
}

type State int
const (
	NOT_LOGGED_IN State = iota
	LOBBY
	IN_CHAT
)
