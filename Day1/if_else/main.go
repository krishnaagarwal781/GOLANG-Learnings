package main

import "fmt"

func main() {
	age := 101

	if age >= 18 && age <= 100 {
		fmt.Println("Karle drive")
	} else if age < 18 {
		fmt.Println("Not a valid person")
	} else {
		fmt.Println("budha gaya hai")
	}

	// can declare variable inside if contruct
	if age2 := 15; age2 <= 18 {
		fmt.Println("bacha hai ")
	}

	// go does not have ternary operator and we have to use normal if else
}
