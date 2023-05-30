package main

import "testing"

func BenchmarkPalindromeOwnSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("A man, a plan, a canal: Panama")
	}
}

func BenchmarkPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2("A man, a plan, a canal: Panama")
	}
}
