package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"client/utils"
)

func main() {

	// Get gateway server host, ip
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	config := utils.GetConfig(data)
	apiGateway, ok := config["api-gateway"].(map[string]interface{});
	if !ok {
		panic(fmt.Errorf("api-gateway config is not a map[string]interface{}"))
	}
	host, ok := apiGateway["host"].(string)
	if !ok {
		panic(fmt.Errorf("api-gateway host is not a string"))
	}
	port, ok := apiGateway["port"].(int)
	if !ok {
		panic(fmt.Errorf("api-gateway port is not an int"))
	}

	// Test connection
	url := "https://" + host + ":" + strconv.Itoa(port)
	resp, err := http.Get(url + "/api/test")
	if err != nil || resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("failed to connect to the gateway server"))
	}

	// Log in
	var userId, pw string
	client := NewDefaultClient(host, port)
	scanner := bufio.NewScanner(os.Stdin)
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

		err = client.Login(userId, pw)
		if err == nil {
			break
		}
		fmt.Println("Failed to login! Try again.")
	}

	quit := false
	inChat := false

	var cmd string
	fmt.Println("Enter '/help' to see available commands.")

	// Execution loop
	for {
		// In chat
		if inChat {
			fmt.Printf("---------- In Chat: %s ----------\n", client.chatName)
			go client.ListenAndDisplay()
			for {
				scanner.Scan()
				m := scanner.Text()
				if m == "/leave" {
					inChat = false
					break
				}
				if err := client.SendMessage(m); err != nil {
					inChat = false
					break
				}
			}
			fmt.Printf("---------- Leave Chat: %s ----------\n", client.chatName)
			client.Leave()
		}

		// Main
		fmt.Printf("(%s)>> ", client.userId)
		scanner.Scan()
		cmd = scanner.Text()
		switch cmd {
		case "/create":
			fmt.Print("Enter chat name: ")
			scanner.Scan()
			chatName := scanner.Text()
			fmt.Print("Enter chat password: ")
			scanner.Scan()
			chatPw := scanner.Text()
			if err := client.Create(chatName, chatPw); err != nil {
				fmt.Println(fmt.Errorf("failed to create chat"))
			}
		case "/delete":
			fmt.Print("Enter chat name to delete: ")
			scanner.Scan()
			chatName := scanner.Text()
			fmt.Print("Enter chat password: ")
			scanner.Scan()
			chatPw := scanner.Text()
			client.Delete(chatName, chatPw)
			fmt.Println("Delete!")
		case "/join":
			fmt.Println("Join!")
			fmt.Print("Enter chat name to join: ")
			scanner.Scan()
			chatName := scanner.Text()
			fmt.Print("Enter chat password: ")
			scanner.Scan()
			chatPw := scanner.Text()
			if err := client.Join(chatName, chatPw); err != nil {
				fmt.Println(err)
			}
			inChat = true
		case "/quit":
			fmt.Println("Quit!")
			quit = true
		}

		if quit {
			break
		}
	}
	fmt.Println("Quit term-chat client")
}
