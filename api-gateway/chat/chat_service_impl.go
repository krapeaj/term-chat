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
	chatRoom := NewChatRoom(chatName, password, []*Client{}, make(chan []byte), s.logger)
	s.chatRooms[chatName] = chatRoom
	// listens to any messages from clients to broadcast
	go chatRoom.Listen()
	return nil
}

func (s *ChatServiceImpl) DeleteChat(userId, chatId string) error {
	s.logger.Printf("Deleting chat '%s' for '%s'\n", chatId, userId)
	return nil
}

func (s *ChatServiceImpl) JoinChat(chatName, password string, client *Client) error {
	if chatName == "" || password == "" {
		return fmt.Errorf("invalid arguemnts '%s', '%s'\n", chatName, password)
	}
	chatRoom, ok := s.chatRooms[chatName]
	if !ok {
		return fmt.Errorf("chat name '%s' does not exist", chatName)
	}
	if chatRoom.password != password {
		return fmt.Errorf("wrong password")
	}
	chatRoom.AddClient(client)

	go client.Listen(chatRoom.ch)
	chatRoom.ch <- []byte(client.UserId + " has joined!")
	return nil
}
