package main

import (
	"testing"
)

func BenchmarkFibIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIterative(30)
	}
}

func BenchmarkFibRecusrive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursive(30)
	}
}

func BenchmarkFibMemoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibMemoization(30)
	}
}
