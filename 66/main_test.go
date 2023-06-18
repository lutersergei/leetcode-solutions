package main

import "testing"

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
		sl2 := plusOne(sl)
		_ = plusOne(sl2)
	}
}

func BenchmarkPlusOneAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl2 := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
		sl3 := plusOneAppend(sl2)
		_ = plusOneAppend(sl3)
	}
}

func BenchmarkPlusOneNeet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl2 := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
		sl3 := plusOneNeetcode(sl2)
		_ = plusOneNeetcode(sl3)
	}
}
