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
	HELP Command = "/help"
	CREATE Command = "/create"
	DELETE Command = "/delete"
	JOIN Command = "/join"
	LEAVE Command = "/leave"
	QUIT Command = "/quit"
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
	var userId, pw string
	for {
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
		if err == nil {
			break
		}
		fmt.Println("Failed to login! Try again.")
	}

	// Execution loop
	quit := false
	inChat := false
	var cmd Command
	fmt.Println("Enter '/help' to see available commands.")
	for !quit {
		// Main
		fmt.Printf("(%s)>> ", client.userId)
		scanner.Scan()
		cmd = Command(scanner.Text())
		switch cmd {

		case HELP:
			printHelp()

		case CREATE:
			if err := createChat(scanner, client); err != nil {
				fmt.Print("ERROR: failed to create chat. ")
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
	fmt.Println("'/create' - Create a chat room.")
	fmt.Println("'/delete' - Delete a chat room.")
	fmt.Println("'/join' - Join a chat room.")
	fmt.Println("'/quit' - Quit client.")

	fmt.Println("Commands in chat:")
	fmt.Println("'/leave' - Leave the current chat room.")
	fmt.Println("'/quit' - Quit client.")
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
