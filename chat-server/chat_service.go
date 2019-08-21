package main

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

func (cs *ChatService) Create(userId string, accessible []string) (string, error) {


}
