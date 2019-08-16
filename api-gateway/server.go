package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// initialize router
	r := mux.NewRouter()

	logger := log.New(os.Stdout, "api-gateway: ", 0)

	var chatService *ChatService
	chatService = NewChatServiceImpl(logger)

	handler := NewHandler(chatService, nil, logger)

	handler.ServeHTTP()


	// TODO:
	// 1) Right now I will have one redis server running, but if I were to make the application scalable,
	// I need to have some sort of service discovery.
	// 2) Maybe later I will use WS for message broadcasting
	// 3) Connection timeout for chat
}
