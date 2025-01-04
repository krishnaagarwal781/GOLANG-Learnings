package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums = []int{1, 2, 3, 4}

	fmt.Println(nums[0:2]) // last wala include nahi hota
	fmt.Println(nums[:3])  // last wala include nahi hota
	fmt.Println(nums[2:])  // last wala include nahi hota

	var nums2 = []int{1, 2, 3, 4}

	fmt.Println(slices.Equal(nums, nums2))

	var two_d = [][]int{{1, 2, 3}, {2, 1}}

	fmt.Println((two_d))
}
