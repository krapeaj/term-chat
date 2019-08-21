package auth

import (
	"api-gateway/cache"
	"api-gateway/persistence"
	"api-gateway/auth/jwt"
	"api-gateway/model"
	"fmt"
	"github.com/google/uuid"
	"log"
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
	// get user from persistence
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
	err = as.cacheService.HMSet(key, val);
	if err != nil {
		return "", err
	}
	return sessionId, err
}

func (as *AuthService) Logout(userId, sessionId string) error {
	return nil
}

func (as *AuthService) GetUser(sessionId string) (*model.User, error) {
	userId := as.cacheService.GetUserIdFromSession(sessionId)
	if userId == "" {
		return nil, fmt.Errorf("user session '%s' does not exist", sessionId)
	}
	user := as.sqlService.GetUser(userId)
	if user == nil {
		return nil, fmt.Errorf("could not find userId '%s' from the database", userId)
	}
	return user, nil
}
