package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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

func (h *Handler) ServeHTTP(addr string, port int) {
	r := mux.NewRouter()

	r.HandleFunc("/say-hello", h.sayHello()).Methods("GET")
	r.HandleFunc("/login", h.login()).Methods("POST")
	r.HandleFunc("/logout", h.logout()).Methods("POST")
	r.HandleFunc("/chat", h.createChat()).Methods("PUT")
	r.HandleFunc("/chat", h.deleteChat()).Methods("DELETE")
	r.HandleFunc("/chat/{chatId}", h.enterChat()).Methods("GET")
	r.HandleFunc("/chat/{chatId}", h.leaveChat()).Methods("DELETE")
	r.HandleFunc("/chat/{chatId}", h.sendMessage()).Methods("POST")

	address := addr + ":" + strconv.Itoa(port)
	server := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	h.logger.Printf("Started listening to port %d\n", port)
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
