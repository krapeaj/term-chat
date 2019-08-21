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

func NewRedisService(client *redis.Client, l *log.Logger) *RedisService {
	return &RedisService{
		client: client,
		logger: l,
	}
}

func (rs *RedisService) HGet(key string) string {
	result := rs.client.HGet(key, "userId")
	if result.Err() != nil {
		return ""
	}
	return result.Val()
}

func (rs *RedisService) HMSet(key string, val map[string]interface{}) error {
	pipe := rs.client.Pipeline()
	pipe.HMSet(key, val)
	rs.client.Expire(key, 30*time.Minute)
	err := pipe.Close()
	if err != nil {
		return err
	}
	return nil
}


