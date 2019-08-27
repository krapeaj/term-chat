package main

import (
	"api-gateway/auth"
	"api-gateway/auth/jwt"
	"api-gateway/cache/redis"
	"api-gateway/chat"
	"api-gateway/persistence"
	_ "database/sql"
	r "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func main() {
	// Get environment variables
	//DEPLOYMENT_TYPE := os.Getenv("DEPLOYMENT_TYPE")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PORT := os.Getenv("REDIS_PORT")
	REDIS_PASSWORD := ""
	REDIS_DB := 0

	// DB - open connection and perform migrations
	db, err := gorm.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Logger
	logger := log.New(os.Stdout, "", 0)

	// SQL service
	sqlService := persistence.NewSQLService(db, logger)

	// JWT service
	jwtService := jwt.NewJwtService(logger)

	// Redis client
	rc := r.NewClient(&r.Options{
		Addr:     REDIS_HOST+":" + REDIS_PORT,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})

	// CacheService
	cacheService := redis.NewRedisService(rc, logger)

	// AuthService
	authService := auth.NewAuthService(cacheService, sqlService, jwtService, logger)

	// ChatService
	chatService := chat.NewChatServiceImpl(cacheService, logger)

	// Create Handler
	handler := NewHandler(chatService, authService, logger)

	// Serve HTTP
	handler.ServeHTTPS()
}
