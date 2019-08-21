package redis

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisService struct {
	client *redis.Client
	logger *log.Logger
}

const (
	userSessionKeyPrefix = "user-session"
)

func NewRedisService(client *redis.Client, l *log.Logger) *RedisService {
	return &RedisService{
		client: client,
		logger: l,
	}
}

func (rs *RedisService) Publish(chatId, userId string) error {
	return nil
}

func (rs *RedisService) Subscribe(chatId, userId string) error {
	return nil
}

func (rs *RedisService) Unsubscribe(chatId, userId string) error {
	return nil
}
