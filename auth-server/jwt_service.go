package main

import "log"

type JwtService struct {
	logger *log.Logger
}

func NewJWTService(l *log.Logger) *JwtService {
	return &JwtService{
		logger: l,
	}
}

func createJWT() string {
	return ""
}

func (js *JwtService) GetJWT(userId string) string {
	return ""
}
