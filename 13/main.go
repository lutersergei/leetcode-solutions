package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	letters := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	length := len(s)
	var ans int
	for i := 0; i < length; i++ {
		if i+1 < length && letters[s[i]] < letters[s[i+1]] {
			ans -= letters[s[i]]
		} else {
			ans += letters[s[i]]
		}
	}
	return ans
}

func romanToIntSwitch(s string) int {
	var sum int

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i < len(s)-1 {
				if s[i+1] == 'V' {
					sum += 4
					i++
					continue
				}
				if s[i+1] == 'X' {
					sum += 9
					i++
					continue
				}
			}
			sum += 1
		case 'V':
			sum += 5
		case 'X':
			if i < len(s)-1 {
				if s[i+1] == 'L' {
					sum += 40
					i++
					continue
				}
				if s[i+1] == 'C' {
					sum += 90
					i++
					continue
				}
			}
			sum += 10
		case 'L':
			sum += 50
		case 'C':
			if i < len(s)-1 {
				if s[i+1] == 'D' {
					sum += 400
					i++
					continue
				}
				if s[i+1] == 'M' {
					sum += 900
					i++
					continue
				}
			}
			sum += 100
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		default:
			panic("wrong roman symbol")
		}
	}

	return sum
}
