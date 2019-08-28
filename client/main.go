package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Command string

const (
	HELP   Command = "/help"
	LOGIN  Command = "/login"
	LOGOUT Command = "/logout"
	SIGNUP Command = "/signup"
	CREATE Command = "/create"
	DELETE Command = "/delete"
	JOIN   Command = "/join"
	LEAVE  Command = "/leave"
	QUIT   Command = "/quit"
)

func main() {

	// Get remote server configuration from file
	serverAddr := getServerAddr()

	// Test connection
	if err := testConnection(serverAddr); err != nil {
		panic(err)
	}

	// Init client
	client := NewDefaultClient(serverAddr)

	// Log in
	scanner := bufio.NewScanner(os.Stdin)

	// Execution loop
	quit := false
	inChat := false
	var cmd Command
	fmt.Println("Enter '/help' to see available commands.")
	for !quit {
		// Main
		fmt.Printf("\n(%s)>> ", client.userId)
		scanner.Scan()
		cmd = Command(scanner.Text())
		switch cmd {

		case HELP:
			printHelp()

		case SIGNUP:
			if err := signup(scanner, client); err != nil {
				fmt.Println("ERROR: failed to sign up.")
				fmt.Println(err)
			}
		case LOGIN:
			if err := login(scanner, client); err != nil {
				fmt.Println("ERROR: failed to log in.")
				fmt.Println(err)
			}
		case LOGOUT:
			if err := logout(client); err != nil {
				fmt.Println("ERROR: failed to log out.")
				fmt.Println(err)
			}
		case CREATE:
			if err := createChat(scanner, client); err != nil {
				fmt.Println("ERROR: failed to create chat.")
				fmt.Println(err)
			}

		case DELETE:
			if err := deleteChat(scanner, client); err != nil {
				fmt.Println("ERROR: failed to delete chat.")
				fmt.Println(err)
			}

		case JOIN:
			if err := joinChat(scanner, client); err != nil {
				fmt.Print("ERROR: failed to join chat. ")
				fmt.Println(err)
			} else {
				inChat = true
			}

		case QUIT:
			quit = true
			break
		}

		// In chat
		if inChat {
			enterChat(scanner, client)
			inChat = false
		}
	}

	fmt.Println("Logging out..")
	client.Logout()
	fmt.Println("Quitting term-chat client..")
}

func getServerAddr() string {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	config := GetConfig(data)
	apiGateway, ok := config["chat-server"].(map[string]interface{});
	if !ok {
		panic(fmt.Errorf("server config is not a map[string]interface{}"))
	}
	host, ok := apiGateway["host"].(string)
	if !ok || host == "" {
		panic(fmt.Errorf("invalid server host"))
	}
	port, ok := apiGateway["port"].(int)
	if !ok {
		panic(fmt.Errorf("invalid server port"))
	}
	return host + ":" + strconv.Itoa(port)
}

func testConnection(serverAddr string) error {
	resp, err := http.Get("https://" + serverAddr + "/api/test")
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to connect to server")
	}
	return nil
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("'/help' - Print this help message.")
	fmt.Println("'/login' - Log in.")
	fmt.Println("'/logout' - Log out.")
	fmt.Println("'/signup' - Sign up.")
	fmt.Println("'/create' - Create a chat room.")
	fmt.Println("'/delete' - Delete a chat room.")
	fmt.Println("'/join' - Join a chat room.")
	fmt.Println("'/quit' - Quit client.")
	fmt.Println()

	fmt.Println("Commands in chat:")
	fmt.Println("'/leave' - Leave the current chat room.")
}

func createChat(scanner *bufio.Scanner, client *Client) error {
	fmt.Print("Enter chat name: ")
	scanner.Scan()
	chatName := scanner.Text()
	fmt.Print("Enter chat password: ")
	scanner.Scan()
	chatPw := scanner.Text()
	return client.Create(chatName, chatPw)
}

func deleteChat(scanner *bufio.Scanner, client *Client) error {
	fmt.Print("Enter chat name to delete: ")
	scanner.Scan()
	chatName := scanner.Text()
	fmt.Print("Enter chat password: ")
	scanner.Scan()
	chatPw := scanner.Text()
	return client.Delete(chatName, chatPw)
}

func joinChat(scanner *bufio.Scanner, client *Client) error {
	fmt.Print("Enter chat name to join: ")
	scanner.Scan()
	chatName := scanner.Text()
	fmt.Print("Enter chat password: ")
	scanner.Scan()
	chatPw := scanner.Text()
	return client.Join(chatName, chatPw)
}

func enterChat(scanner *bufio.Scanner, client *Client) {
	fmt.Printf("---------- In Chat: %s ----------\n", client.chatName)
	go client.ListenAndDisplay()
	for {
		scanner.Scan()
		m := scanner.Text()
		if Command(m) == LEAVE {
			break
		}
		if err := client.SendMessage(m); err != nil {
			break
		}
	}
	client.Leave()
	fmt.Printf("---------- Leave Chat: %s ----------\n", client.chatName)
}

func login(scanner *bufio.Scanner, client *Client) error {
	if client.sessionId != "" {
		return fmt.Errorf("user already logged in")
	}
	var userId, pw string
	for {
		fmt.Print("User ID: ")
		scanner.Scan()
		userId = scanner.Text()
		if userId != "" {
			break
		}
		fmt.Println("User ID cannot be empty!")
	}
	for {
		fmt.Print("User PW: ")
		scanner.Scan()
		pw = scanner.Text()
		if pw != "" {
			break
		}
		fmt.Println("Password cannot be empty!")
	}
	err := client.Login(userId, pw)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func logout(client *Client) error {
	if client.sessionId == "" {
		return fmt.Errorf("user not logged in")
	}
	err := client.Logout()
	if err != nil {
		return err
	}
	return nil
}

func signup(scanner *bufio.Scanner, client *Client) error {
	if client.sessionId != "" {
		return fmt.Errorf("user is logged in")
	}
	fmt.Print("Enter User ID to sign up with: ")
	scanner.Scan()
	userId := scanner.Text()
	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()
	return client.Signup(userId, password)
}
