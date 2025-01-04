package main

import "fmt"

func main() {
	var num []int
	fmt.Println(num == nil)
	fmt.Println(len(num))

	var num2 = make([]int, 2) //-> (type of slice, initial size of slice)
	fmt.Println(cap(num2))    // capacity batata hai ye - which is maximum no. of elementscan fit

	var num3 = make([]int, 4, 5) //  iska initial size toh 2 hai but iski capacity 5 bhi hai
	fmt.Println(num3)            // [0,0]
	fmt.Println(cap(num3))       // returns 5

	// How to add element in slice

	num3 = append(num3, 1)
	num3 = append(num3, 1)
	num3 = append(num3, 1)
	num3 = append(num3, 1)
	num3 = append(num3, 1)
	fmt.Println(num3)
	fmt.Println(len(num3)) // ye array ka current length batata hai
	fmt.Println(cap(num3)) // ye array ka total current size batata hai

	//new way to create slice
	nums4 := []int{}
	nums4 = append(nums4, 1)

	num3[4] = 4
	fmt.Println(num3)
	fmt.Println(len(nums4))
	fmt.Println(cap(nums4))

}
