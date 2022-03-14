package main

import "testing"

func BenchmarkGenerateAlertID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateAlertID()
	}
}

func BenchmarkGoogleID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoogleGen()
	}
}
