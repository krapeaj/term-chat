package service

import (
	"chat-server/cache"
	"log"
)

type SessionService struct {
	cacheService cache.CacheService
	logger       *log.Logger
}

func NewSessionService(cs cache.CacheService, l *log.Logger) *SessionService {
	return &SessionService{
		cacheService: cs,
		logger:       l,
	}
}
