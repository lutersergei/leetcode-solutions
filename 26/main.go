package main

import "fmt"

// Remove Duplicates from Sorted Array
func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//arr := []int{1, 1, 2, 3}
	k := removeDuplicates(arr)
	fmt.Println(k, arr)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	var i, j = 0, 1
	var prevDupl bool

	for j < len(nums) {
		if nums[i] == nums[j] {
			prevDupl = true
			j++
			continue
		}

		if prevDupl {
			nums[i+1] = nums[j]
			i++
			prevDupl = false
			continue
		}

		j++
		i++
		prevDupl = false
	}

	return i + 1
}

func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	k, i := 1, 1
	for i < length {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
		i++
	}
	return k
}
