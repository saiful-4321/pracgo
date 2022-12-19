package main
import "fmt"

func main() {
	fmt.Println("Wellcome to loop section")

	days := []string{"Sunday","Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Printf("Days are: %v\n", days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("index is : %v, and value is: %v \n",index, value)
	// }

	roughValue := 1
	for roughValue < 10 {
		if(roughValue == 5) {
			goto lco
		}
		// if(roughValue == 5) {
		// 	break
		// }
		if(roughValue == 5) {
			roughValue++
			continue
		}
		fmt.Println("value is :", roughValue)
		roughValue++
	}

	lco:
		fmt.Println("here to jump")
}