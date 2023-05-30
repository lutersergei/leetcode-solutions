package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("1739!(#_ "))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {
	var res int
	var tmp int

	ab := [1024]int{}

	for i := 0; i < len(s); i++ {
		//fmt.Println(string(s[i]))
		if ab[s[i]] == 0 {
			tmp++
			ab[s[i]] = tmp
		} else {
			ab[s[i]] = 0
		}

		if tmp > res {
			res = tmp
		}
	}

	return res
}
