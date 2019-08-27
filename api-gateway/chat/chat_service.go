package chat

type ChatService interface {
	CreateChat(chatName, password string) error
	DeleteChat(userId, chatId string) error
	JoinChat(chatName, password string, client *Client) error
}
