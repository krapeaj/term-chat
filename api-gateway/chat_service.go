package main

type ChatService interface {
	CreateChat(userId string) error
	DeleteChat(userId, chatId string) error
	EnterChat(chatId, password, userId string) error
	LeaveChat(chatId, userId string) error
	SendMessage(userId, chatId, message string) error
}
