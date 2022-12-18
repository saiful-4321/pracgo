package main

import "fmt"

func main() {
	fmt.Println("welcome to map section")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "Paython"
	languages["RB"] = "Ruby"

	fmt.Println("List of all languages:",languages)
	fmt.Println("Js shorts for:",languages["js"])
	
	delete(languages, "RB")
	fmt.Println("List of all languages:",languages)

	// loops in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}