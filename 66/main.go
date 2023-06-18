package main

import "fmt"

func main() {
	fmt.Println(plusOneAppend([]int{1, 2, 3}))
	fmt.Println(plusOneAppend([]int{9, 9, 9}))
	fmt.Println(plusOneAppend([]int{9}))
}

func plusOne(digits []int) []int {
	r := len(digits) - 1
	var reminder = 1
	for r >= 0 {
		num := digits[r] + reminder

		if num == 10 {
			reminder = 1
			num = 0
		} else {
			reminder = 0
		}

		digits[r] = num
		r--
	}

	// for first element
	if reminder == 1 {
		tmp := make([]int, len(digits)+1)
		tmp[0] = 1
		copy(tmp[1:], digits)
		digits = tmp
	}

	return digits
}

func plusOneAppend(digits []int) []int {
	r := len(digits) - 1
	var reminder = 1
	for r >= 0 {
		num := digits[r] + reminder

		if num == 10 {
			reminder = 1
			num = 0
		} else {
			reminder = 0
		}

		digits[r] = num
		r--
	}

	// for first element
	if reminder == 1 {
		digits = append(digits, digits[len(digits)-1])
		digits[0] = 1
	}

	return digits
}

func plusOneNeetcode(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
