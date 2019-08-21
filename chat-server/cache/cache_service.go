package cache

type CacheService interface {
	Publish(chatId, userId string) error
	Subscribe(chatId, userId string) error
	Unsubscribe(chatId, userId string) error
}
