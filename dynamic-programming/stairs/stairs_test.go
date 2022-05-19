package main

import "testing"

func Benchmark_tabulation_stairs(b *testing.B) {
	n := 1000
	
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tabulation_getminstairs(n)
	}
}

func Benchmark_memoization_stairs(b *testing.B) {
	n := 1000
	memo := make([]int,n+1)
	
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoization_getminstairs(n,memo)
	}
}

func Benchmark_recursive_stairs(b *testing.B) {
	n := 1000
	
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recursive_getminstairs(n)
	}
}