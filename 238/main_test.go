package main

import "testing"

func BenchmarkProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		productExceptSelf([]int{1, 2, 3, 4})
	}
}
