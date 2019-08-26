package chat

import (
	"api-gateway/cache"
	"fmt"
	"log"
)

type ChatServiceImpl struct {
	cacheService cache.CacheService
	logger       *log.Logger
	chatRooms    map[string]*ChatRoom
}

const (
	chatSessionKeyPrefix = "chat-session"
)

func NewChatServiceImpl(cs cache.CacheService, l *log.Logger) *ChatServiceImpl {
	return &ChatServiceImpl{
		cacheService: cs,
		logger:       l,
		chatRooms:    make(map[string]*ChatRoom),
	}
}

func (s *ChatServiceImpl) CreateChat(chatName, password string) error {
	s.logger.Printf("Creating chat '%s'\n", chatName)
	if chatName == "" || password == "" {
		return fmt.Errorf("invalid arguemnts '%s', '%s'\n", chatName, password)
	}
	chatRoom := NewChatRoom(chatName, []*Client{}, s.logger)
	s.chatRooms[chatName] = chatRoom
	return nil
}

func (s *ChatServiceImpl) DeleteChat(userId, chatId string) error {
	s.logger.Printf("Deleting chat '%s' for '%s'\n", chatId, userId)
	return nil
}

func (s *ChatServiceImpl) JoinChat(chatName, password  string, client *Client) error {
	if chatName == "" || password == "" {
		return fmt.Errorf("invalid arguemnts '%s', '%s'\n", chatName, password)
	}
	chatRoom, ok := s.chatRooms[chatName]
	if !ok {
		return fmt.Errorf("chat name '%s' does not exist", chatName)
	}
	chatRoom.AddClient(client)
	return nil
}

func (s *ChatServiceImpl) LeaveChat(chatId, userId string) error {
	s.logger.Printf("User '%s' leaving chat '%s'\n", userId, chatId)
	return nil
}

func (s *ChatServiceImpl) SendMessage(userId, chatName, message string) error {
	if userId == "" || chatName == "" {
		return fmt.Errorf("invalid arguemnts '%s', '%s'\n", userId, chatName)
	}
	s.logger.Printf("User '%s' sending message in chat '%s'\n", userId, chatName)
	chatRoom, ok := s.chatRooms[chatName]
	if !ok {
		return fmt.Errorf("no chat room '%s'", chatName)
	}
	chatRoom.BroadcastMessage(message)
	return nil
}
