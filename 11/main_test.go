package main

import "testing"

func BenchmarkArea2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxArea2(h)
	}
}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxArea(h)
	}
}
