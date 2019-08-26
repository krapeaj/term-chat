package chat

import (
	"log"
)

type ChatRoom struct {
	chatName string
	password string
	clients  []*Client
	logger   *log.Logger
}

func NewChatRoom(chatName string, clients []*Client, l *log.Logger) *ChatRoom {
	return &ChatRoom{
		chatName: chatName,
		clients:  clients,
		logger:   l,
	}
}

func (cr *ChatRoom) AddClient(client *Client) {
	cr.clients = append(cr.clients, client)
}

func (cr *ChatRoom) BroadcastMessage(message string) {
	for _, c := range cr.clients {
		if err := c.BroadcastMessage(message); err != nil {
			cr.logger.Println(err)
		}
	}
}
