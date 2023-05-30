package main

import "fmt"

// Container With Most Water
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 8, 6, 200, 200, 4, 8, 3, 7}))
}

var h = []int{1, 8, 6, 200, 200, 4, 8, 3, 7}

func maxArea(height []int) int {
	var res, currArea int
	l, r := 0, len(height)-1

	for l != r {
		currArea = (r - l) * min(height[l], height[r])
		res = max(res, currArea)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

func maxArea2(height []int) int {
	n := len(height)
	l, r := 0, n-1

	result := 0
	for l <= r {
		area := (r - l) * min(height[l], height[r])
		result = max(result, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
