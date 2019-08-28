package chat

type ChatService interface {
	CreateChat(chatName, password string) error
	DeleteChat(chatName, password string) error
	JoinChat(chat *ChatRoom, client *Client) error
	GetChatToJoin(chatName, password string) (*ChatRoom, bool)
}
