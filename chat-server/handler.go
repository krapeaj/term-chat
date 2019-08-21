package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	chatService *ChatService
	logger      *log.Logger
}

func NewHandler(cs *ChatService, l *log.Logger) *Handler {
	return &Handler{
		chatService: cs,
		logger:      l,
	}
}

func (h *Handler) ServeHTTP(addr string) {
	r := mux.NewRouter()

	r.HandleFunc("/chat", h.createChat()).Methods("PUT")
	r.HandleFunc("/chat/{chatId}", h.deleteChat()).Methods("DELETE")
	r.HandleFunc("/chat/{chatId}", h.enterChat()).Methods("GET")

	h.logger.Println("Starting server at " + addr + " ...")
	h.logger.Fatal(http.ListenAndServe(addr, r))
}

func (h *Handler) createChat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// parse body
		var b []byte
		_, err := r.Body.Read(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var body map[string]interface{}
		err = json.Unmarshal(b, body)
		if err != nil {
			h.logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		userId, ok := body["userId"].(string)
		if !ok {
			h.logger.Println(fmt.Errorf("invalid body"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		accessible, ok := body["accessible"].([]string)
		if !ok {
			h.logger.Println(fmt.Errorf("invalid body"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// create chat session
		chatId, err := h.chatService.Create(userId, accessible)
		if err != nil {
			h.logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("chatId", chatId)
	}
}

func (h *Handler) deleteChat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) enterChat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
