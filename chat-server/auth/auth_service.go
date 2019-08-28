package auth

import (
	"chat-server/cache"
	"chat-server/persistence"
	"chat-server/auth/jwt"
	"chat-server/model"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

type AuthService struct {
	cacheService cache.CacheService
	jwtService   *jwt.JwtService
	sqlService   *persistence.SQLService
	logger       *log.Logger
}

const (
	userSessionKeyPrefix = "user-session"
)

func NewAuthService(cs cache.CacheService, ss *persistence.SQLService, js *jwt.JwtService, l *log.Logger) *AuthService {
	return &AuthService{
		cacheService: cs,
		sqlService:   ss,
		jwtService:   js,
		logger:       l,
	}
}

func (as *AuthService) Login(userId, password string) (string, error) {
	if userId == "" || password == "" {
		return "", fmt.Errorf("invalid userId or password")
	}
	// get user from DB
	user := as.sqlService.GetUser(userId)
	if !user.IsPasswordMatch(password) {
		return "", fmt.Errorf("passwords do not match")
	}

	// create session
	uid, err := uuid.NewRandom()
	sessionId := uid.String()
	if err != nil {
		return "", err
	}

	key := userSessionKeyPrefix + ":" + sessionId
	val := map[string]interface{}{
		"userId": userId,
	}
	as.logger.Println("session key: " + key)
	err = as.cacheService.HMSet(key, val);
	if err != nil {
		return "", err
	}
	return sessionId, err
}

func (as *AuthService) Logout(sessionId string) (string, error) {
	key := userSessionKeyPrefix + ":" + sessionId
	userId, err := as.cacheService.Del(key)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (as *AuthService) GetUser(sessionId string) (*model.User, error) {
	key := userSessionKeyPrefix + ":" + sessionId
	userId := as.cacheService.HGet(key)
	if userId == "" {
		return nil, fmt.Errorf("user session '%s' does not exist", sessionId)
	}
	user := as.sqlService.GetUser(userId)
	if user == nil {
		return nil, fmt.Errorf("could not find userId '%s' from the database", userId)
	}
	return user, nil
}

func (as *AuthService) CreateUser(userId, password string) (*model.User, error) {
	if userId == "" || password == "" {
		return nil, fmt.Errorf("empty userId or password")
	}
	user := &model.User{
		UserId: userId,
		Password: password,
		DateCreated: time.Now().String(),
	}
	err := as.sqlService.SetUser(user)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return user, nil
}
