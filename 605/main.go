package main

import "fmt"

// Can Place Flowers
func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	var nextBad bool

	for i, flow := range flowerbed {
		if flow == 1 {
			nextBad = true
			continue
		}

		if nextBad == false && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			nextBad = true
		} else {
			nextBad = false
		}

		if n == 0 {
			return true
		}
	}

	return false
}
