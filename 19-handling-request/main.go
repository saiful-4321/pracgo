package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://agentapi.paywellonline.com/HealthCheck/SystemHealthCheck/healthCheckStatus"

func main() {
	fmt.Println("Welcome to handle request section")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is %T \n", response)
	// fmt.Println(response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)

	fmt.Println(content)
}
