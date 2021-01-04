package main

import "testing"

var result string

func BenchmarkRandStringBytesRmndr(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = RandStringBytesRmndr(12)
	}
	result = r
}
