package main

import (
	"testing"
)

func Benchmark_fibonacci(b *testing.B) {
	b.ReportAllocs()

	for i:=0; i<b.N;i++{
		fibonacci(25)
	}
}

func Benchmark_memoization(b *testing.B) {
	b.ReportAllocs()

	for i:=0; i<b.N;i++{
		memoization(25)
	}
}

func Benchmark_tabulation(b *testing.B) {
	b.ReportAllocs()

	for i:=0; i<b.N;i++{
		tabulation(25)
	}
}

func Benchmark_fibDynamic(b *testing.B) {
	b.ReportAllocs()

	for i:=0; i<b.N;i++{
		fibDynamic(25)
	}
}
