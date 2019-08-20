package cache

type CacheService interface {
	GetUser(sessionId string) string
	SetUser(userId, sessionId string) error

	Publish(chatId, userId string) error
	Subscribe(chatId, userId string) error
	Unsubscribe(chatId, userId string) error
}
