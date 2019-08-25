package redis

import (
	"fmt"
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
	pipe.Expire(key, 30*time.Minute)
	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	err = pipe.Close()
	if err != nil {
		return err
	}
	return nil
}

func (rs *RedisService) Del(key string) (string, error) {
	strCmd := rs.client.HGet(key, "userId")
	if strCmd.Err() != nil {
		rs.logger.Println(fmt.Errorf("session key '%s' does not exist", key))
		return "", strCmd.Err()
	}
	intCmd := rs.client.Del(key)
	if intCmd.Err() != nil {
		rs.logger.Println(fmt.Errorf("unexpected error while removing key '%s'", key))
		return "", intCmd.Err()
	}
	return strCmd.Val(), nil
}
