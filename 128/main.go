package main

import (
	"fmt"
	"math"
)

/*
Longest Consecutive Sequence
*/
func main() {
	//fmt.Println(longestConsecutive([]int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999}))
	fmt.Println(longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}))
}

func longestConsecutive(nums []int) int {
	var resultSeq = 0

	if len(nums) == 0 {
		return resultSeq
	}

	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		set[num] = struct{}{}
	}

	var maxSeq int
	for num := range set {
		if _, ok := set[num-1]; !ok {
			// start sequence if left number not in set
			maxSeq = 1
			for {
				if _, ok = set[num+maxSeq]; ok {
					maxSeq++
				} else {
					if maxSeq > resultSeq {
						resultSeq = maxSeq
					}
					break
				}
			}
		}

	}

	return resultSeq
}

/*
bruteforce version
*/
func longestConsecutiveBrut(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var min = math.MaxInt32
	var max = math.MinInt32

	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		m[num] = struct{}{}
	}

	var seq = 1
	var tmp = 1

	for i := min + 1; i <= max; i++ {
		if _, ok := m[i]; ok {
			tmp++
			if tmp > seq {
				seq = tmp
			}
		} else {
			tmp = 0
		}
	}

	return seq
}
