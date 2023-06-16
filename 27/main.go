package main

import "fmt"

// 27. Remove Element
func main() {
	//sl := []int{3, 2, 2, 3}
	sl := []int{2}

	fmt.Println(removeElement(sl, 2), sl)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var i, j int

	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	return i
}
