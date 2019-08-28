package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	serverAddr string
	userId     string
	sessionId  string
	chatName   string
	ws         *websocket.Conn
}

const (
	HTTPS           = "https://"
	WSS             = "wss://"
	API_PREFIX      = "/api"
	ENDPOINT_LOGIN  = API_PREFIX + "/login"
	ENDPOINT_LOGOUT = API_PREFIX + "/logout"
	ENDPOINT_CREATE = API_PREFIX + "/chat"
	ENDPOINT_DELETE = API_PREFIX + "/chat"
	WEBSOCKET       = "/websocket"
)

func NewDefaultClient(serverAddr string) *Client {
	return &Client{
		serverAddr: serverAddr,
	}
}

func (c *Client) Login(userId, password string) error {
	fmt.Printf("Trying to log in as '%s'...\n", userId)

	req, err := http.NewRequest("POST", HTTPS+c.serverAddr+ENDPOINT_LOGIN, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(userId, password)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return fmt.Errorf("log in rejected by server")
	}
	if resp.StatusCode == 200 {
		sessionId := resp.Header.Get("Set-Cookie")
		if sessionId == "" {
			return fmt.Errorf("received no cookie")
		}
		c.sessionId = sessionId
		c.userId = userId
		fmt.Println("Log in successful!")
		return nil
	}
	return fmt.Errorf("unexpected log in failure")
}

func (c *Client) Logout() error {
	if c.sessionId == "" {
		return fmt.Errorf("user is not logged in")
	}
	fmt.Println("Logging out...")

	req, err := http.NewRequest("POST", HTTPS+c.serverAddr+ENDPOINT_LOGOUT, bytes.NewBuffer(nil))
	if err != nil {
		return err
	}
	req.Header.Add("session-id", c.sessionId)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		fmt.Println("Log out successful!")
		return nil
	}
	if resp.StatusCode == 400 {
		return fmt.Errorf("session has ended already")
	}
	return fmt.Errorf("failed to log out")
}

func (c *Client) Create(chatName, password string) error {
	fmt.Printf("Creating chat room '%s'...\n", chatName)
	req, _ := http.NewRequest("PUT", HTTPS+c.serverAddr+ENDPOINT_CREATE, nil)
	req.Header.Add("session-id", c.sessionId)
	req.Header.Add("chat-name", chatName)
	req.Header.Add("password", password)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create chat room")
	}

	fmt.Printf("Successfully created chat room '%s'.\n", chatName)
	return nil
}

func (c *Client) Delete(chatName, chatPw string) error {
	fmt.Printf("Deleting chat '%s'...\n", chatName)
	body, err := json.Marshal(map[string]interface{}{
		"chatName": chatName,
		"password": chatPw,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", HTTPS+c.serverAddr+ENDPOINT_LOGOUT, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("session-id", c.sessionId)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Successfully delete chat room '%s'\n", chatName)
		return nil
	}
	return fmt.Errorf("failed to delete chat room '%s'", chatName)
}

func (c *Client) Join(chatName, chatPw string) error {
	fmt.Printf("Joining chat '%s'...\n", chatName)
	header := http.Header{}
	header.Add("session-id", c.sessionId)
	header.Add("chat-name", chatName)
	header.Add("password", chatPw)
	conn, resp, err := websocket.DefaultDialer.Dial(WSS+c.serverAddr+WEBSOCKET, header)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to establish websocket connection"))
		if resp.StatusCode == http.StatusForbidden {
			return fmt.Errorf("wrong chat name or password")
		}
		return fmt.Errorf("failed to join chat '%s'\n", chatName)
	}
	c.ws = conn
	c.chatName = chatName
	fmt.Printf("Successfully joined chat '%s'!\n", chatName)
	return nil
}

func (c *Client) Leave() {
	c.ws.WriteMessage(websocket.CloseMessage, []byte(c.userId+" leaving chat"))
	c.ws.Close()
	c.chatName = ""
}

func (c *Client) SendMessage(message string) error {
	if err := c.ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		fmt.Println(fmt.Errorf("failed to send message to server"))
		return err
	}
	return nil
}

func (c *Client) ListenAndDisplay() {
	for {
		t, msg, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(msg))
		if t == websocket.CloseMessage {
			break
		}
	}
	c.ws.Close()
}
