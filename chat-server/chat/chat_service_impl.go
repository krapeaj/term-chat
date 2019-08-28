package chat

import (
	"chat-server/cache"
	"fmt"
	"log"
)

type ChatServiceImpl struct {
	cacheService cache.CacheService
	logger       *log.Logger
	chatRooms    map[string]*ChatRoom
}

const (
	chatSessionKeyPrefix = "chat-session"
)

func NewChatServiceImpl(cs cache.CacheService, l *log.Logger) *ChatServiceImpl {
	return &ChatServiceImpl{
		cacheService: cs,
		logger:       l,
		chatRooms:    make(map[string]*ChatRoom),
	}
}

func (s *ChatServiceImpl) CreateChat(chatName, password string) error {
	s.logger.Printf("Creating chat '%s'\n", chatName)
	if chatName == "" || password == "" {
		return fmt.Errorf("invalid arguemnts '%s', '%s'\n", chatName, password)
	}
	chatRoom := NewChatRoom(chatName, password, []*Client{}, make(chan []byte), s.logger)
	s.chatRooms[chatName] = chatRoom
	// listens to any messages from clients to broadcast
	go chatRoom.Listen()
	return nil
}

func (s *ChatServiceImpl) DeleteChat(chatName, password string) error {
	c, ok := s.chatRooms[chatName]
	if !ok {
		return fmt.Errorf("no chat named '%s'", chatName)
	}

	if c.password != password {
		return fmt.Errorf("wrong password")
	}
	for _, client := range c.clients {
		client.Disconnect("chat room deleted")
	}
	delete(s.chatRooms, chatName)
	return nil
}

func (s *ChatServiceImpl) JoinChat(chatRoom *ChatRoom, client *Client) error {
	if chatRoom == nil {
		return fmt.Errorf("invalid chat room")
	}
	chatRoom.AddClient(client)
	go client.Listen(chatRoom.ch)
	chatRoom.ch <- []byte(client.UserId + " has joined!")
	return nil
}

func (s *ChatServiceImpl) GetChatToJoin(chatName, password string) (*ChatRoom, bool) {
	chat, ok := s.chatRooms[chatName]
	if !ok {
		return nil, false
	}
	if chat.password != password {
		return chat, false
	}
	return chat, true
}
