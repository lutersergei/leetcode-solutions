package main

import (
	"fmt"
	"sort"
)

// Two Sum II - Input Array Is Sorted
func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15, 16, 17}, 13))
	//fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}, 542))
}

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	for i, number := range numbers {
		if ind, flag := findInSlice(numbers[i+1:], target-number); flag == true {
			res[0] = i + 1
			res[1] = ind + i + 2
			return res
		}
	}

	return res
}

func twoSum2(numbers []int, target int) []int {
	res := make([]int, 2)
	for i, number := range numbers {
		if ind, flag := search(numbers[i+1:], target-number); flag == true {
			res[0] = i + 1
			res[1] = ind + i + 2
			return res
		}
	}

	return res
}

func twoSum3(numbers []int, target int) []int {
	res := make([]int, 2)
	for i, number := range numbers {
		if ind, flag := searchB(numbers[i+1:], target-number); flag == true {
			res[0] = i + 1
			res[1] = ind + i + 2
			return res
		}
	}

	return res
}

func twoSum2P(numbers []int, target int) []int {
	p1 := 0
	p2 := len(numbers) - 1

	for p1 < p2 {
		n := numbers[p1] + numbers[p2]
		if n == target {
			return []int{p1 + 1, p2 + 1}
		}

		if n > target {
			p2--
		} else {
			p1++
		}
	}

	return []int{0, 0}
}

func searchB(nums []int, target int) (int, bool) {
	idx := sort.SearchInts(nums, target)
	if idx >= len(nums) || nums[idx] != target {
		return -1, false
	}
	return idx, true
}

func findInSlice(sl []int, target int) (int, bool) {
	if len(sl) == 0 {
		return -1, false
	}
	if len(sl) == 1 {
		if sl[0] == target {
			return 0, true
		}
		return -1, false
	}
	mid := len(sl) / 2
	if sl[mid] == target {
		return mid, true
	}
	if sl[mid] > target {
		return findInSlice(sl[:mid], target)
	}
	ind, flag := findInSlice(sl[mid+1:], target)
	return ind + mid + 1, flag
}

func search(nums []int, target int) (int, bool) {
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid, true
		}
	}

	return -1, false
}
