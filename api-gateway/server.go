package main

import (
	"fmt"
	"log"
	"os"

	"api-gateway/auth"
	"api-gateway/cache/redis"
	"api-gateway/auth/jwt"
	"api-gateway/persistence"

	r "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func main() {
	// logger
	logger := log.New(os.Stdout, "api-gateway: ", 0)

	// DB connection

	/* Lower abstractions */
	// jwt service
	jwtService := jwt.NewJwtService(logger)
	// sql service
	db, err := gorm.Open("mysql", "test.persistence")
	if err != nil {
		panic(fmt.Errorf("failed to open a DB connection"))
	}
	sqlService := persistence.NewSQLService(db, logger)
	// user session
	us := r.NewClient(&r.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})
	userSessionService := redis.NewRedisService(us, logger)
	// chat session
	cs := r.NewClient(&r.Options{
		Addr:     "0.0.0.0:7379",
		Password: "",
		DB:       0,
	})
	chatSessionService := redis.NewRedisService(cs, logger)


	/* Higher abstractions */
	// auth service
	authService := auth.NewAuthService(userSessionService, sqlService, jwtService, logger)
	// chat service
	chatService := NewChatServiceImpl(chatSessionService, logger)

	// create handler
	handler := NewHandler(chatService, authService, logger)

	// serve HTTP
	handler.ServeHTTPS()

	// TODO:
	// 1) Right now I will have one redis server running, but if I were to make the application scalable,
	// I need to have some sort of service discovery.
	// 2) Maybe later I will use WS for message broadcasting
	// 3) Connection timeout for chat
}
