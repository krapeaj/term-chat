package main

type Client interface {
	Login(userId, password string) error
	Logout() error
	Create(chatName, password string) error
	Delete(chatName, password string) error
	Join(chatName, password string) error
	Leave()
	SendMessage(message string) error
}
