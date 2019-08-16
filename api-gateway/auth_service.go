package main

type AuthService interface {
	login(userId, password string) error
	logout(userId, sessionId string) error
}
