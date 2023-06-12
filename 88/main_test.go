package main

import "testing"

var a = []int{1, 2, 3, 8, 0, 0, 0}
var c = []int{2, 5, 6}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge(a, 4, c, 3)
	}
}

func BenchmarkMerge2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge2(a, 4, c, 3)
	}
}
