package main

import (
	"chat-server/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	sessionService *service.SessionService
	chatService    *service.ChatService
	logger         *log.Logger
}

func NewHandler(ss *service.SessionService, cs *service.ChatService, l *log.Logger) *Handler {
	return &Handler{
		sessionService: ss,
		chatService:    cs,
		logger:         l,
	}
}

func (h *Handler) ServeHTTP(addr string) {
	r := mux.NewRouter()

	r.HandleFunc("/user/session", h.getUserSession()).Methods("GET")
	r.HandleFunc("/user/session", h.createUserSession()).Methods("PUT")
	r.HandleFunc("/user/session", h.deleteUserSession()).Methods("DELETE")

	h.logger.Fatal(http.ListenAndServe(addr, r))
}

func (h *Handler) getUserSession() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("getUserSession"))
	}
}

func (h *Handler) createUserSession() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("createUserSession"))
	}
}

func (h *Handler) deleteUserSession() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("deleteUserSession"))
	}
}
