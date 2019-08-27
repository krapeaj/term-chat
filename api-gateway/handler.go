package main

import (
	"api-gateway/auth"
	"api-gateway/chat"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

	r.HandleFunc("/test", h.test()).Methods("GET")
	r.HandleFunc("/login", h.login()).Methods("POST")
	r.HandleFunc("/logout", h.logout()).Methods("POST")
	r.HandleFunc("/chat", h.createChat()).Methods("PUT")
	r.HandleFunc("/chat", h.deleteChat()).Methods("DELETE")
	r.HandleFunc("/websocket", h.joinChat()).Methods("GET")

	server := &http.Server{
		Addr:         "0.0.0.0:433",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	h.logger.Println("Will start listening on port 433")
	log.Fatal(server.ListenAndServeTLS( "/go/bin/app/certs/server.crt", "/go/bin/app/certs/server.key"))
}

func (h *Handler) test() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			h.logger.Println(fmt.Errorf("failed to write in body"))
		}
		w.WriteHeader(http.StatusOK)
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
		sessionId := r.Header.Get("session-id")
		if sessionId == "" {
			h.logger.Println(fmt.Errorf("no session-id"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userId, err := h.authService.Logout(sessionId)
		if err != nil {
			h.logger.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		h.logger.Printf("Logout successful for user '%s'\n", userId)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) createChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("session-id")
		if sessionId == "" {
			h.logger.Println("error: no sessionId")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// check user session
		_, err := h.authService.GetUser(sessionId)
		if err != nil {
			h.logger.Printf("error: failed to get user with sessionId '%s'\n", sessionId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// get chat info
		chatName := r.Header.Get("chat-name")
		password := r.Header.Get("password")
		if chatName == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// create chat - spins up go routine to listen for messages to broadcast
		err = h.chatService.CreateChat(chatName, password)
		if err != nil {
			h.logger.Println(fmt.Errorf("failed to create chat"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h.logger.Printf("Successfully created chat '%s'\n", chatName)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *Handler) deleteChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("DELETE CHAT")
	}
}

func (h *Handler) joinChat() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("session-id")
		if sessionId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// get user from session
		u, err := h.authService.GetUser(sessionId)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// get chat info
		chatName := r.Header.Get("chat-name")
		password := r.Header.Get("password")

		// create client
		client := chat.NewClient(u.UserId, h.logger)
		// upgrade to ws
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		err = client.AddWebSocketConn(conn)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// spins up a go routine for listening to client
		err = h.chatService.JoinChat(chatName, password, client)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
