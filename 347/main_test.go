package main

import "testing"

func BenchmarkTopK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topKFrequent([]int{-1, -1, 1, 2, 2, 3, 4, 4, 4}, 3)
		topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	}
}

func BenchmarkTopK2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topKFrequent2([]int{-1, -1, 1, 2, 2, 3, 4, 4, 4}, 3)
		topKFrequent2([]int{1, 1, 1, 2, 2, 3}, 2)
	}
}
