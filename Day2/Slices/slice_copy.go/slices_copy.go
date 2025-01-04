package main

import "fmt"

func main() {
	var nums = make([]int, 2, 5)
	nums = append(nums, 1)
	var nums2 = make([]int, len(nums), 5)

	fmt.Println(nums)
	fmt.Println(nums2)

	// how to copy

	copy(nums2, nums) // nums is copied to nums2
	fmt.Println(nums2, nums)
}
