package service

import (
	"chat-server/cache"
	"log"
)

type ChatService struct {
	cacheService cache.CacheService
	logger       *log.Logger
}

func NewChatService(cs cache.CacheService, l *log.Logger) *ChatService {
	return &ChatService{
		cacheService: cs,
		logger:       l,
	}
}
