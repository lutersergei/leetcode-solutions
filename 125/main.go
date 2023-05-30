package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

// Valid Palindrome
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("    "))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = clearString(s)

	l := len(s)
	if l < 2 {
		return true
	}
	mid := l / 2

	for i := 0; i <= mid; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}

	return true
}

func isPalindrome2(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i = i + 1
		j = j - 1
	}

	return true
}
