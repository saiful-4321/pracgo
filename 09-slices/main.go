package main 

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice learning")

	var fruitlist = []string{"Apple", "Malta", "Banana"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)
	
	fruitlist = append(fruitlist, "Mango", "Jambura")
	fmt.Println(fruitlist)
	
	fruitlist = append(fruitlist[:1])
	fmt.Println(fruitlist)
	
	highscore := make([]int, 4)
	highscore[0] = 123
	highscore[1] = 112
	highscore[2] = 212
	highscore[3] = 1122
	// highscore[4] = 777

	// append used to re allocation memory
	highscore = append(highscore, 333, 444, 555)

	fmt.Println(highscore)
	
	
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))
}