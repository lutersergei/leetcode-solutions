package main

import "testing"

var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var k int

func BenchmarkRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k = removeDuplicates(arr)
	}
}

func BenchmarkRemove2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k = removeDuplicates2(arr)
	}
}
