package main

import (
	"testing"
)

func BenchmarkConstant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		returnConst()
	}
}

func BenchmarkVariable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		returnVar()
	}
}
