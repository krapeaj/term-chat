package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	UserId string
	logger *log.Logger
	ws     *websocket.Conn
}

func NewClient(userId string, l *log.Logger) *Client {
	return &Client{
		UserId: userId,
		logger: l,
	}
}

func (c *Client) AddWebSocketConn(conn *websocket.Conn) error {
	if c.ws != nil {
		return fmt.Errorf("web socket connection already exists")
	}
	c.ws = conn
	return nil
}

func (c *Client) Listen(ch chan<- []byte) {
	for {
		t, m, err := c.ws.ReadMessage()
		if err != nil || t == websocket.CloseMessage {
			ch <- []byte(c.UserId + " has left.")
			break
		}
		message := c.UserId + ": " + string(m)
		ch <- []byte(message)
	}
	c.logger.Println("client has been disconnected")
	c.ws.Close()
}

func (c *Client) SendMessage(message []byte) {
	if c.ws == nil {
		c.logger.Println(fmt.Errorf("web socket connection does not exist"))
		return
	}
	err := c.ws.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		c.logger.Println(fmt.Errorf("failed to send message to user '%s'", c.UserId))
		c.logger.Println(err)
		c.ws.Close()
	}
}

func (c *Client) Disconnect(reason string) {
	if c.ws == nil {
		c.logger.Println(fmt.Errorf("web socket connection does not exist"))
		return
	}
	c.ws.WriteMessage(websocket.CloseMessage, []byte(reason))
	c.ws.Close()
}
