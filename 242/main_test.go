package main

import "testing"

func BenchmarkAnagramMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagramMy("anagram", "nagaram")
		isAnagramMy("cat", "tac")
	}
}

func BenchmarkAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram("anagram", "nagaram")
		isAnagram("cat", "tac")
	}
}
