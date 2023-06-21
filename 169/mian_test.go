package main

import "testing"

var sl = []int{2, 2, 1, 1, 1, 2, 2}

func BenchmarkMajority(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = majorityElement(sl)
	}
}

func BenchmarkMajoritySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = majorityElementSort(sl)
	}
}
