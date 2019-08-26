package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	UserId   string
	logger   *log.Logger
	ws       *websocket.Conn
}

func NewClient(userId string, l *log.Logger) *Client {
	return &Client{
		UserId:   userId,
		logger:   l,
	}
}

func (c *Client) AddWebSocketConn(conn *websocket.Conn) error {
	if c.ws != nil {
		return fmt.Errorf("web socket connection already exists")
	}
	c.ws = conn
	return nil
}

func (c *Client) BroadcastMessage(message string) error {
	if c.ws == nil {
		return fmt.Errorf("web socket connection does not exist")
	}
	err := c.ws.WriteJSON(map[string]string{
		"userId":  c.UserId,
		"message": message,
	})
	if err != nil {
		c.logger.Println(fmt.Errorf("failed to broadcast message to user '%s'", c.UserId))
	}
	return nil
}
