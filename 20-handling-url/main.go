package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=react&id=10"

func main() {
	fmt.Println("Welcome to URL section")

	fmt.Println(myurl)

	// parsing url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// depth in query param
	qparams := result.Query()
	fmt.Printf("The type of the query params is : %T \n", qparams)

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=Saiful",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
