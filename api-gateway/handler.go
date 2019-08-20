package main

import (
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

func (h *Handler) ServeHTTPS() {
	r := mux.NewRouter()

	r.HandleFunc("/say-hello", h.sayHello()).Methods("GET")
	r.HandleFunc("/login", h.login()).Methods("POST")
	r.HandleFunc("/logout", h.logout()).Methods("POST")
	r.HandleFunc("/chat", h.createChat()).Methods("PUT")
	r.HandleFunc("/chat", h.deleteChat()).Methods("DELETE")
	r.HandleFunc("/chat/{chatId}", h.enterChat()).Methods("GET")
	r.HandleFunc("/chat/{chatId}", h.leaveChat()).Methods("DELETE")
	r.HandleFunc("/chat/{chatId}", h.sendMessage()).Methods("POST")

	server := &http.Server{
		Addr:         ":433",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	h.logger.Println("Will start listening on port 433")
	log.Fatal(server.ListenAndServeTLS( "/go/bin/app/certs/server.crt", "/go/bin/app/certs/server.key"))
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
		userId, _, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Printf("User '%s' attempting to log in.", userId)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) logout() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("LOGOUT")
	}
}

func (h *Handler) createChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := http.StatusOK
		userId := r.Header.Get("userId")
		if userId == "" {
			h.logger.Println("Error: requires userId.")
			result = http.StatusBadRequest
		} else {
			err := h.chatService.CreateChat(userId)
			if err != nil {
				h.logger.Println("Error: failed to create chat.")
				result = http.StatusBadRequest
			}
		}
		w.WriteHeader(result)
	}
}

func (h *Handler) deleteChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("DELETE CHAT")
	}
}

func (h *Handler) sendMessage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SEND MESSAGE")
	}
}

func (h *Handler) enterChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ENTER CHAT")
	}
}

func (h *Handler) leaveChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("LEAVE CHAT")
	}
}
