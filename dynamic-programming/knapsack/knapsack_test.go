package main

import "testing"

func Benchmark_tabulation_knapsack(b *testing.B) {
	capacity := 50
	value := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	n := 3
	
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tabulation_knapsack(weight,value,capacity,n)
	}
}

func Benchmark_recursive_knapsack(b *testing.B) {
	capacity := 50
	value := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	n := 3
	
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recursive_knapsack(weight,value,capacity,n)
	}
}

func Benchmark_memoization_knapsack(b *testing.B) {
	capacity := 50
	value := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	n := 3
	
	table := make([][]int, 1002)
	for i := range table {
		table[i] = make([]int, 1002)
	}
	for i := 0; i < 1002; i++ {
		for j := 0; j < 1002; j++ {
			table[i][j] = -1
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoization_knapsack(weight,value,capacity,n,table)
	}
}