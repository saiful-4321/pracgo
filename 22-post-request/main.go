package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to post request section")

	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json data
	requestBody := strings.NewReader(`
		"coursename" : "javascript",
		"courseid" : 10,
		"price" : 100
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
