package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to handle get request section")

	PerformGetRequestTwo()
}

func PerformGetRequest() {
	const weburl = "http://localhost:8000/get"

	response, err := http.Get(weburl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Stateus code is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformGetRequestTwo() {
	const weburl = "http://localhost:8000/get"

	response, err := http.Get(weburl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Stateus code is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is:", byteCount)
	fmt.Println(responseString.String())
}
