package chat

import (
	"api-gateway/cache"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type ChatServiceImpl struct {
	cacheService cache.CacheService
	logger       *log.Logger
}

const (
	chatSessionKeyPrefix = "chat-session"
)

func NewChatServiceImpl(cs cache.CacheService, l *log.Logger) *ChatServiceImpl {
	return &ChatServiceImpl{
		cacheService: cs,
		logger:       l,
	}
}

func (s *ChatServiceImpl) CreateChat(userId, chatName, password string) (string, error) {
	s.logger.Printf("'%s' creating chat '%s'\n", userId, chatName)
	if userId == "" || chatName == "" || password == "" {
		return "", fmt.Errorf("invalid arguemnts '%s', '%s', '%s'\n", userId, chatName, password)
	}
	rand, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	chatId := rand.String()

	// TODO: PUB/SUB?

	return chatId, nil
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
