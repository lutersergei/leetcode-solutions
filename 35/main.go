package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(sl, 0))
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := ((high - low) / 2) + low
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
