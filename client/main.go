package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"client/utils"
)

func main() {

	// get gateway server host, ip
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


	// ask for session to gateway server
	url := "http://" + host + ":" + strconv.Itoa(port)

	resp, err := http.Get(url + "/api/say-hello")
	if err != nil {
		panic(fmt.Errorf("failed to connect to the gateway server"))
	}
	fmt.Println(resp.Header)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body")
	}
	fmt.Println(string(b))

	client := NewDefaultClient("https://" + url)
	err = client.Login("krapeaj", "krapeaj")
	if err != nil {
		fmt.Println("failed to login")
	}
	//quit := false
	/*
	for {

		if quit {
			break
		}
	}
	*/
}
