package main

import "testing"

func BenchmarkMyVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternately("shdfkjsdklfjskdgkjsdjf;sldfsdjflksdf", "shdfkjsdklfjskdgkjsdjf;sldfsdjflksdf")
	}
}

func BenchmarkAlter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternately2("shdfkjsdklfjskdgkjsdjf;sldfsdjflksdf", "shdfkjsdklfjskdgkjsdjf;sldfsdjflksdf")
	}
}

func BenchmarkAlter2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeAlternately3("shdfkjsdklfjskdgkjsdjf;sldfsdjflksdf", "shdfkjsdklfjskdgkjsdjf;sldfsdjflksdf")
	}
}
