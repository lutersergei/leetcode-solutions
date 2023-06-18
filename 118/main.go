package main

import "fmt"

// Pascal's Triangle
func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(6))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		sl := make([]int, i+1)

		sl[0], sl[i] = 1, 1

		for j := 1; j < len(sl)-1; j++ {
			sl[j] = res[i-1][j] + res[i-1][j-1]
		}
		res[i] = sl
	}

	return res
}
