package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

type Handler struct {
	chatService *ChatService
	authService *AuthService
	logger      *log.Logger
}

func NewHandler(cs *ChatService, as *AuthService, l *log.Logger) *Handler {
	return &Handler{
		chatService: cs,
		authService: as,
		logger: l,
	}
}

func (h *Handler) ServeHTTP() {
	r := mux.NewRouter()

	r.HandleFunc("/api/say-hello", sayHello).Methods("GET")
	r.HandleFunc("/api/login", login).Methods("GET")
	r.HandleFunc("/api/logout", logout).Methods("GET")
	r.HandleFunc("/api/chat", createChat).Methods("POST")
	r.HandleFunc("/api/chat", deleteChat).Methods("DELETE")
	r.HandleFunc("/api/chat/{chatId}", enterChat).Methods("POST")
	r.HandleFunc("/api/chat/{chatId}", leaveChat).Methods("DELETE")
	r.HandleFunc("/api/chat/{chatId}", sendMessage).Methods("POST")

	server := &http.Server{
		Handler: r,
		Addr: "127.0.0.1.8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func (h *Handler) sayHello(w http.ResponseWriter, r *http.Request) {
	return 
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		fmt.Println("Error!")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
}

func logout(w http.ResponseWriter, r *http.Request) {
}

func createChat(w http.ResponseWriter, r *http.Request) {
	var userId string
	result := http.StatusOK
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		h.logger.Println("Error: failed to parse request body.")
		result = http.StatusBadRequest
	}
	err = h.chatService.createChat(userId)
	if err != nil {
		h.logger.Println("Error: failed to create chat.")
		result = http.StatusBadRequest
	}
	w.WriteHeader(result)
}

func (h *Handler) deleteChat(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) sendMessage(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) enterChat(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) leaveChat(w http.ResponseWriter, r *http.Request) {
}
