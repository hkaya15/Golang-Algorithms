package main

import (
	"fmt"
	"math"
)

// https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/

func main() {
	capacity := 50
	value := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	n := 3
	
	fmt.Println(tabulation_knapsack(weight, value, capacity, n))

	table := make([][]int, 1002)
	for i := range table {
		table[i] = make([]int, 1002)
	}
	for i := 0; i < 1002; i++ {
		for j := 0; j < 1002; j++ {
			table[i][j] = -1
		}
	}

	fmt.Println(memoization_knapsack(weight, value, capacity, n, table))
}


// Benchmark_tabulation_knapsack-8          1576297               719.8 ns/op          1760 B/op          5 allocs/op
func tabulation_knapsack(weight, value []int, capacity, n int) int {
	table := make([][]int, n+1)
	
	for i := 0; i <= n; i++ {
		table[i] = make([]int, capacity+1)
	}
	
	for i := 0; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if weight[i-1] <= j {
				table[i][j] = int(math.Max(float64(value[i-1]+table[i-1][j-weight[i-1]]), float64(table[i-1][j]))) // Accept the value as weight is acceptable.
			} else {
				table[i][j] = table[i-1][j] // Ignore the value is exceeded.
			}
		}
	}
	return table[n][capacity]
}


// Benchmark_recursive_knapsack-8          23444539                50.70 ns/op            0 B/op          0 allocs/op
func recursive_knapsack(weight, value []int, capacity, n int) int {

	if n == 0 || capacity == 0 {
		return 0
	}

	if weight[n-1] <= capacity {
		return int(math.Max(float64(value[n-1]+recursive_knapsack(weight, value, capacity-weight[n-1], n-1)), float64(recursive_knapsack(weight, value, capacity, n-1))))
	} else if weight[n-1] > capacity { 
		return recursive_knapsack(weight, value, capacity, n-1)
	}

	return 1
}

// Benchmark_memoization_knapsack-8        578471229                2.064 ns/op           0 B/op          0 allocs/op
func memoization_knapsack(weight, value []int, capacity, n int, table [][]int) int {
	if capacity == 0 || n == 0 {
		return 0
	}
	if table[capacity][n] != -1 {
		return table[capacity][n]
	}
	if weight[n-1] <= capacity {
		table[capacity][n] = int(math.Max(float64(value[n-1]+memoization_knapsack(weight, value, capacity-weight[n-1], n-1, table)), float64(memoization_knapsack(weight, value, capacity, n-1, table))))
		return table[capacity][n]
	} else if weight[n-1] > capacity {
		table[capacity][n] = memoization_knapsack(weight, value, capacity, n-1, table)
		return table[capacity][n]
	}
	return 1
}

