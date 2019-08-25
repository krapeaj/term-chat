package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DefaultClient struct {
	serverAddr string
	userId     string
	password   string
	sessionId  string
	chatId     string
	state      State
}

const (
	ENDPOINT_API_PREFIX = "/api"
	ENDPOINT_LOGIN      = ENDPOINT_API_PREFIX + "/login"
	ENDPOINT_LOGOUT     = ENDPOINT_API_PREFIX + "/logout"
	ENDPOINT_CREATE     = ENDPOINT_API_PREFIX + "/chat"
	ENDPOINT_DELETE     = ENDPOINT_API_PREFIX + "/chat/%s"
	ENDPOINT_ENTER      = ENDPOINT_API_PREFIX + "/chat/%s"
	ENDPOINT_LEAVE      = ENDPOINT_API_PREFIX + "/chat/%s/leave"
	ENDPOINT_MESSAGE    = ENDPOINT_API_PREFIX + "/chat/%s"
)

func NewDefaultClient(serverAddr string) *DefaultClient {
	return &DefaultClient{serverAddr: serverAddr}
}

func (c *DefaultClient) Login(userId, password string) error {
	fmt.Printf("Trying to log in as '%s'...\n", userId)

	req, err := http.NewRequest("POST", c.serverAddr+ENDPOINT_LOGIN, nil)
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
		fmt.Println("Log in successful!")
		return nil
	}
	return fmt.Errorf("unexpected log in failure")
}

func (c *DefaultClient) Logout() error {
	if c.sessionId == "" {
		return fmt.Errorf("user is not logged in")
	}
	fmt.Println("Logging out...")
	req, err := http.NewRequest("POST", c.serverAddr+ENDPOINT_LOGOUT, nil)
	if err != nil {
		return err
	}
	c.setCookie(req)
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

func (c *DefaultClient) Create(chatName, password string) error {
	fmt.Printf("Creating chat room '%s'...\n", chatName)
	body, err := json.Marshal(map[string]interface{}{
		"chatName": chatName,
		"password": password,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", c.serverAddr+ENDPOINT_CREATE, bytes.NewBuffer(body))
	c.setCookie(req)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 201 {
		chatId := resp.Header.Get("chat-id")
		if chatId == "" {
			return fmt.Errorf("did not get chat-id in response")
		}
		c.chatId = chatId
		fmt.Printf("Successfully created chat room '%s'.\n", chatName)
		return nil
	}
	return fmt.Errorf("failed to create chat room")
}

func (c *DefaultClient) Delete(chatName string) error {
	fmt.Printf("Deleting chat '%s'...\n", chatName)
	return nil
}

func (c *DefaultClient) Enter(chatName string) error {
	fmt.Printf("Entering chat '%s'...\n", chatName)
	return nil
}

func (c *DefaultClient) Leave() error {
	fmt.Println("Leaving chat...")
	return nil
}

func (c *DefaultClient) SendMessage(message string) error {
	fmt.Println("Message sent??")
	return nil
}

func (c *DefaultClient) setCookie(req *http.Request) error {
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: c.sessionId,
	}
	req.AddCookie(cookie)
	return nil
}
