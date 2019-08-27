package main

type Client interface {
	Login(userId, password string) error
	Logout() error
	Create(chatName, password string) error
	Delete(chatName, password string) error
	Join(chatName, password string) error
	Leave() error
	SendMessage(message string)
}

type State int
const (
	NOT_LOGGED_IN State = iota
	LOBBY
	IN_CHAT
)
