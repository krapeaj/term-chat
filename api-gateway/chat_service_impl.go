package main

import "log"

type ChatServiceImpl struct {
	logger *log.Logger
}

func NewChatServiceImpl(l *log.Logger) *ChatServiceImpl {
	return &ChatServiceImpl{
		logger: l,
	}
}

func (s *ChatServiceImpl) CreateChat(userId string) error {
	s.logger.Printf("Creating chat for '%s'\n", userId)
	return nil
}

func (s *ChatServiceImpl) DeleteChat(userId, chatId string) error {
	s.logger.Printf("Deleting chat '%s' for '%s'\n", chatId, userId)
	return nil
}

func (s *ChatServiceImpl) EnterChat(chatId, password, userId string) error {
	s.logger.Printf("User '%s' entering chat '%s'\n", userId, chatId)
	return nil
}

func (s *ChatServiceImpl) LeaveChat(chatId, userId string) error {
	s.logger.Printf("User '%s' leaving chat '%s'\n", userId, chatId)
	return nil
}

func (s *ChatServiceImpl) SendMessage(userId, chatId, message string) error {
	s.logger.Printf("User '%s' sending message in chat '%s'\n", userId, chatId)
	return nil
}
