package main

import (
	"fmt"
	"math"
)

func main() {
	// n:=10000000 stack overflow
	n := 50000
	memo := make([]int, n+1)
	fmt.Println(recursive_getminstairs(n))
	fmt.Println(memoization_getminstairs(n, memo))
	fmt.Println(tabulation_getminstairs(n))

}

// Benchmark_tabulation_stairs-8             114266             10492 ns/op            8192 B/op          1 allocs/op
func tabulation_getminstairs(n int) int {
	table := make([]int, n+1)

	for i := 2; i <= n; i++ {
		if (i%2 == 0) && i%3 != 0 {
			table[i] = 1 + int(math.Min(float64(table[i-1]), float64(table[i/2])))
		} else if (i%3 == 0) && i%2 != 0 {

			table[i] = 1 + int(math.Min(float64(table[i-1]), float64(table[i/3])))

		} else if (i%2 == 0) && (i%3 == 0) {
			table[i] = 1 + int(math.Min(float64(table[i-1]), math.Min(float64(table[i/2]), float64(table[i/3]))))
		} else {
			table[i] = 1 + table[i-1]
		}
	}

	return table[n]
}

// Benchmark_memoization_stairs-8          587254495                2.039 ns/op           0 B/op          0 allocs/op
func memoization_getminstairs(n int, memo []int) int {
	if n == 1 {
		return 0
	}

	if memo[n] != 0 {
		return memo[n]
	}

	result := memoization_getminstairs(n-1, memo)

	if n%2 == 0 {
		result = int(math.Min(float64(result), float64(memoization_getminstairs(n/2, memo))))
	}

	if n%3 == 0 {
		result = int(math.Min(float64(result), float64(memoization_getminstairs(n/3, memo))))
	}

	memo[n] = 1 + result
	return memo[n]
}

// Benchmark_recursive_minstairs-8                1        75543187500 ns/op              0 B/op          0 allocs/op
func recursive_getminstairs(n int) int {
	if n == 1 {
		return 0
	}
	result := recursive_getminstairs(n - 1)

	if n%2 == 0 {
		result = int(math.Min(float64(result), float64(recursive_getminstairs(n/2))))
	}

	if n%3 == 0 {
		result = int(math.Min(float64(result), float64(recursive_getminstairs(n/3))))
	}

	return result + 1
}
