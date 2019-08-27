package chat

import (
	"log"
)

type ChatRoom struct {
	chatName string
	password string
	ch       chan []byte
	clients  []*Client
	logger   *log.Logger
}

func NewChatRoom(chatName, password string, clients []*Client, ch chan []byte, l *log.Logger) *ChatRoom {
	return &ChatRoom{
		chatName: chatName,
		password: password,
		ch:       ch,
		clients:  clients,
		logger:   l,
	}
}

func (cr *ChatRoom) AddClient(client *Client) {
	cr.clients = append(cr.clients, client)
}

func (cr *ChatRoom) GetClient(userId string) *Client {
	for _, c := range cr.clients {
		if c.UserId == userId {
			return c
		}
	}
	return nil
}

func (cr *ChatRoom) Listen() {
	// TODO: idle timeout
	for {
		msg, open := <-cr.ch
		if !open {
			break
		}
		cr.BroadcastMessage(msg)
	}
}

func (cr *ChatRoom) RemoveClient(client *Client) {
	for i, c := range cr.clients {
		if c == client {
			cr.clients[i] = cr.clients[len(cr.clients)-1]
			cr.clients = cr.clients[:len(cr.clients)-1]
			return
		}
	}
}

func (cr *ChatRoom) BroadcastMessage(message []byte) {
	for _, c := range cr.clients {
		c.SendMessage(message)
	}
}
