package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study.")

	// to get current date
	presentTime := time.Now()
	fmt.Println(presentTime)

	// 2006 for year
	// 01 for month
	// 02 for date
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// 2020 for year
	// month, day, hour, munites, secound, nano sec, time zone
	createdDate := time.Date(2020, time.January, 11, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15-04-05 Monday"))

}