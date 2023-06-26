package main

func main() {
	println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	println(numIdenticalPairs([]int{1, 1, 1, 1}))
	println(numIdenticalPairs([]int{1, 2, 3}))
}

func numIdenticalPairs(nums []int) int {
	var res int

	for i, num := range nums {
		for ind := i + 1; ind < len(nums); ind++ {
			if nums[ind] == num {
				res++
			}
		}
	}

	return res
}
