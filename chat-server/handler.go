package main

import (
	"chat-server/cache"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	cacheService cache.CacheService
	logger       *log.Logger
}

func NewHandler(cs cache.CacheService, l *log.Logger) *Handler {
	return &Handler{
		cacheService: cs,
		logger: l,
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
