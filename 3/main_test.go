package main

import (
	"reflect"
	"testing"
	"unsafe"
)

var a []byte

func BenchmarkNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a = []byte("test string")
	}
}

func BenchmarkStrTOBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a = strToBytes("test string")
	}
}

func strToBytes(s string) []byte {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	shdr := reflect.SliceHeader{Data: h.Data, Len: h.Len, Cap: h.Len}
	return *(*[]byte)(unsafe.Pointer(&shdr))
}
