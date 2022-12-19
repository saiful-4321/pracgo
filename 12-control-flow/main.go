package main

import "fmt"

func main() {
	fmt.Println("wellcome to control flow")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount == 10 {
		result = "Total 10 users append"
	} else {
		result = "more then regular user"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Number is more then 3")
	} else {
		fmt.Println("Number is less then 3")
	}
}