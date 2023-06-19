package main

import (
	"fmt"
	"sort"
)

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}

type NumArray struct {
	Sl []int
}

func Constructor(nums []int) NumArray {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	var preSum = []int{0}
	preSum = append(preSum, nums...)
	return NumArray{Sl: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Sl[right+1] - this.Sl[left]
}
