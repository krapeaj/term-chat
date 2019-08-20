package redis

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

type RedisService struct {
	client *redis.Client
	logger *log.Logger
}

const (
	userSessionKeyPrefix = "user-session"
)

func NewRedisService(addr, password string, db int, l *log.Logger) *RedisService {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisService{
		client: client,
		logger: l,
	}
}

func (rs *RedisService) GetUser(sessionId string) string {
	key := userSessionKeyPrefix + ":" + sessionId
	result := rs.client.HGet(key, "userId");
	if result.Err() != nil {
		return ""
	}
	return result.Val()
}

func (rs *RedisService) SetUser(userId, sessionId string) error {
	pipe := rs.client.Pipeline()
	key := userSessionKeyPrefix + ":" + sessionId
	val := map[string]interface{} {
		"userId": userId,
	}
	pipe.HMSet(key, val)
	rs.client.Expire(key, 30 * time.Minute)
	err := pipe.Close()
	if err != nil {
		return err
	}
	return nil
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
