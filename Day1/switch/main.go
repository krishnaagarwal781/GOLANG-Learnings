package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch

	i := 5

	switch i {
	case 1:
		fmt.Println("one") // agar i ki value one hai toh ye chalega
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("Other")
	}

	// multi condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Chutti hai bhai ji")

	default:
		fmt.Println("Kaam karo bc")
	}

}
