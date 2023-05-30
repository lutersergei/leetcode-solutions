package main

import (
	"fmt"
)

// Move Zeroes
func main() {
	sl := []int{1, 0, 3, 12} //	[1,3,12,0,0]
	moveZeroes(sl)

	fmt.Println(sl)
}

func moveZeroes(nums []int) {
	var l, r = 0, 0
	for ; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}

//func moveZeroes(nums []int) {
//	var l, r = 0, len(nums) - 1
//
//	for l <= r {
//		if nums[r] == 0 {
//			r--
//			continue
//		}
//
//		if nums[l] != 0 {
//			l++
//			continue
//		}
//
//		nums[l], nums[r] = nums[r], nums[l]
//		l++
//		r--
//	}
//}
