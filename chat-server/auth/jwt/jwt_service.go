package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

type JwtService struct {
	logger *log.Logger
}

type JWT string

func NewJwtService(l *log.Logger) *JwtService {
	return &JwtService{
		logger: l,
	}
}

func (js *JwtService) CreateJWT(userId string, accessible []string) (JWT, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     userId,
		"accessible": accessible,
	})
	token, err := t.SignedString("secret")
	if err != nil {
		return JWT(""), fmt.Errorf("failed to create JWT token")
	}

	js.logger.Printf("Created JWT token: %s\n", token)
	return JWT(token), nil
}
