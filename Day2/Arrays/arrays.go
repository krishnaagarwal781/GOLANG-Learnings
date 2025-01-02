package main

import "fmt"

func main() {
	var nums [5]int
	var bol [5]bool
	var bhaiji [5]string

	nums[0] = 1
	bol[0] = true
	bhaiji[0] = "chacha ji"

	//array ke length ke liye hai
	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)

	fmt.Println(len(bol))
	fmt.Println(bol[0])
	fmt.Println(bol)

	fmt.Println(len(bhaiji))
	fmt.Println(bhaiji[0])
	fmt.Println(bhaiji)

	numing := [4]int{1, 2, 3, 4}

	fmt.Println((numing))

	// 2 d array
	nums2 := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(nums2)
}
