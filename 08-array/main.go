package main

import "fmt"

func main() {
	fmt.Println("Welcome to array learning section")

	var fruitLists [4]string

	fruitLists[0] = "Apple"
	fruitLists[1] = "Tomato"
	fruitLists[3] = "Banana"

	fmt.Println(fruitLists)
	fmt.Println(len(fruitLists))

	var vegList = [3] string{"potato", "beans", "mashroom"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}