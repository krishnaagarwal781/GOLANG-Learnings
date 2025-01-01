package main

import "fmt"

// for is the only construct for looping
func main() {

	// while looping try karte hai
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// infintie loop

	// for {
	// 	fmt.Println(i)
	// }

	//for loop hai ye
	for i := 0; i <= 10; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
		break
	}

	// does not include last number

	for i := range 3 {
		fmt.Println(i)
	}

}
