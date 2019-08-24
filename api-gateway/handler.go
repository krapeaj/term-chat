package main

import (
	"api-gateway/auth"
	"api-gateway/chat"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	chatService chat.ChatService
	authService *auth.AuthService
	logger      *log.Logger
}

func NewHandler(cs chat.ChatService, as *auth.AuthService, l *log.Logger) *Handler {
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
		Addr:         "0.0.0.0:433",
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
		userId, pw, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			h.logger.Println(fmt.Errorf("no credentials provided"))
			return
		}
		h.logger.Printf("User '%s' attempting to log in.\n", userId)

		sessionId, err := h.authService.Login(userId, pw)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			h.logger.Println(err)
			return
		}
		w.Header().Add("Set-Cookie", sessionId)
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
		sessionCookie, err := r.Cookie("sessionId")
		if err != nil {
			h.logger.Println("error: no sessionId")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// get user from session
		sessionId := sessionCookie.Value
		user, err := h.authService.GetUser(sessionId)
		if err != nil {
			h.logger.Printf("error: failed to get user with sessionId '%s'\n", sessionId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var b []byte
		r.Body.Read(b)
		var data map[string]interface{}
		json.Unmarshal(b, data)
		chatName, ok := data["chatName"].(string)
		if !ok {
			h.logger.Println(fmt.Errorf("invalid request parameters"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		password, ok := data["password"].(string)
		if !ok {
			h.logger.Println(fmt.Errorf("invalid request parameters"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// create chat
		chatId, err := h.chatService.CreateChat(user.UserId, chatName, password)
		if err != nil {
			h.logger.Println("Error: failed to create chat.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("CHATID: " + chatId)
		w.WriteHeader(http.StatusOK)
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
