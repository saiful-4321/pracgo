package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to post request section")

	PerformPostFormRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json data
	requestBody := strings.NewReader(`
		{
			"coursename" : "javascript",
			"courseid" : 10,
			"price" : 100
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json generate for request data

	data := url.Values{}

	data.Add("name", "Saiful islam")
	data.Add("email", "bd.saif321@gmail.com")
	data.Add("mobile", "01644331161")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
