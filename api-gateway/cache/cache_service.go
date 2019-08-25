package cache

type CacheService interface {
	HGet(key string) string
	HMSet(key string, val map[string]interface{}) error
	Del(key string) (string, error)
}
