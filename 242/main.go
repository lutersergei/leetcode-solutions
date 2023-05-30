package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("cat", "tac"))
	fmt.Println(isAnagram("cat", "tad"))

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := [26]uint{}

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}

func isAnagramMy(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := makeMap(s)
	m2 := makeMap(t)

	for run, c := range m1 {
		if c2, ok := m2[run]; ok {
			if c != c2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func makeMap(s string) map[rune]int {
	m := make(map[rune]int, 26)

	for _, r := range s {
		m[r]++
	}

	return m
}
