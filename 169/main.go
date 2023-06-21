package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElementSort([]int{3, 2, 3}))
	fmt.Println(majorityElementSort([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	m := make(map[int]int, len(nums))

	mlen := (len(nums) / 2) + 1

	for _, num := range nums {
		m[num]++

		c := m[num]
		if c == mlen {
			return num
		}
	}
	return 0
}

func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
