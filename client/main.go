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
	client := NewDefaultClient(url)
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

	//quit := false
	//for {
	//
	//
	//	if quit {
	//		break
	//	}
	//}
	fmt.Println("Quit term-chat client")
	//quit := false
	/*
	for {

		if quit {
			break
		}
	}
	*/
}
