package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	chatService ChatService
	authService AuthService
	logger      *log.Logger
}

func NewHandler(cs ChatService, as AuthService, l *log.Logger) *Handler {
	return &Handler{
		chatService: cs,
		authService: as,
		logger:      l,
	}
}

func (h *Handler) ServeHTTP() {
	r := mux.NewRouter()

	r.HandleFunc("/api/say-hello", h.sayHello()).Methods("GET")
	r.HandleFunc("/api/login", h.login()).Methods("GET")
	r.HandleFunc("/api/logout", h.logout()).Methods("GET")
	r.HandleFunc("/api/chat", h.createChat()).Methods("POST")
	r.HandleFunc("/api/chat", h.deleteChat()).Methods("DELETE")
	r.HandleFunc("/api/chat/{chatId}", h.enterChat()).Methods("POST")
	r.HandleFunc("/api/chat/{chatId}", h.leaveChat()).Methods("DELETE")
	r.HandleFunc("/api/chat/{chatId}", h.sendMessage()).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func (h *Handler) sayHello() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			fmt.Println("Error!")
		}
	}
}

func (h *Handler) login() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("LOGIN")
	}
}

func (h *Handler) logout() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("LOGOUT")
	}
}

func (h *Handler) createChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var userId string
		result := http.StatusOK
		err := json.NewDecoder(r.Body).Decode(&userId)
		if err != nil {
			h.logger.Println("Error: failed to parse request body.")
			result = http.StatusBadRequest
		}
		err = h.chatService.CreateChat(userId)
		if err != nil {
			h.logger.Println("Error: failed to create chat.")
			result = http.StatusBadRequest
		}
		w.WriteHeader(result)
	}
}

func (h *Handler) deleteChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *Handler) sendMessage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *Handler) enterChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h *Handler) leaveChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
