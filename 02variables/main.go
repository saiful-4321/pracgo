package main

import "fmt"

const status int = 200

func main() {
	var username string = "saiful"
	fmt.Println(username)
	fmt.Printf("%T \n", username)
	
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("%T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("%T \n", smallVal)

	var smallFloat float32 = 255.155555555555
	fmt.Println(smallFloat)
	fmt.Printf("%T \n", smallFloat)

	// default value and some aliases
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("%T \n", defaultInt)

	// implicit type
	var website = "facebook.com"
	fmt.Println(website)
	fmt.Printf("%T \n", website)

	// no var style
	numberOfUser := 20.9
	fmt.Println(numberOfUser)
	fmt.Printf("%T \n", numberOfUser)

	fmt.Println(status)
}