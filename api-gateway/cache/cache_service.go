package cache

type CacheService interface {
	HGet(key string) string
	HMSet(key string, val map[string]interface{}) error
}
