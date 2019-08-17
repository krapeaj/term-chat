package main

import (
	"log"
	"os"
)

func main() {
	// logger
	logger := log.New(os.Stdout, "api-gateway: ", 0)

	// chat service
	chatService := NewChatServiceImpl(logger)

	// create handler
	handler := NewHandler(chatService, nil, logger)

	// serve HTTP
	handler.ServeHTTP("0.0.0.0", 8080)


	// TODO:
	// 1) Right now I will have one redis server running, but if I were to make the application scalable,
	// I need to have some sort of service discovery.
	// 2) Maybe later I will use WS for message broadcasting
	// 3) Connection timeout for chat
}
