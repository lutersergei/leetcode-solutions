package main

import "fmt"

/*
Product of Array Except Self
*/
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, 0, 3}))
}

func productExceptSelf(nums []int) []int {
	var prefix, postfix = 1, 1
	res := make([]int, len(nums))

	for i := range nums {
		res[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
