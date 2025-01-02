package main

import "fmt"

func main() {
	var num []int
	fmt.Println(num == nil)
	fmt.Println(len(num))

	var num2 = make([]int, 2)
	fmt.Println(cap(num2))
}
