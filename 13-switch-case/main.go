package main
import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("Wellcome to switch case part")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6)+1

	fmt.Println(diceNumber)

	switch diceNumber {
		case 1 : 
			fmt.Println("dice Number value is 1")
		case 2 : 
			fmt.Println("dice Number value is 2")
			fallthrough
		case 3 : 
			fmt.Println("dice Number value is 3")
			fallthrough
		case 4 : 
			fmt.Println("dice Number value is 4")
		case 5 : 
			fmt.Println("dice Number value is 5")
		case 6 : 
			fmt.Println("Dice number 6 Vamos!")
		default: 
			fmt.Println("Dhora khawa public")
	}
}