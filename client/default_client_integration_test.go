package main

import "testing"

const (
	localServerAddr = "https://localhost:8000"
)

func TestDefaultClient_LoginAndLogout(t *testing.T) {
	client := NewDefaultClient(localServerAddr)
	userId := "krapeaj"
	pw := "krapeajpw"
	err := client.Login(userId, pw)
	if err != nil {
		t.Errorf("failed to log in")
		return
	}
	t.Log("log in success")

	err = client.Logout()
	if err != nil {
		t.Errorf("failed to log out")
		return
	}
	t.Log("log out success")
}

func TestDefaultClient_Create(t *testing.T) {

}

func TestDefaultClient_Delete(t *testing.T) {

}

func TestDefaultClient_Enter(t *testing.T) {

}

func TestDefaultClient_Leave(t *testing.T) {

}

func TestDefaultClient_SendMessage(t *testing.T) {

}
