package main

import "testing"

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintStyle("lk;sdjggsdfg", "asdfulkvjk")
	}
}

func BenchmarkStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsStyle("lk;sdjggsdfg", "asdfulkvjk")
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsConcat("lk;sdjggsdfg", "asdfulkvjk")
	}
}
