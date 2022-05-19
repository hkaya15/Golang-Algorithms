package main

import "fmt"

// Memoization is an optimization technique used to speed up programs by storing the results of expensive function calls and returning the cached result when the same inputs occur again.
// Tabulation is an approach where you solve a dynamic programming problem by first filling up a table, and then compute the solution to the original problem based on the results in this table.

var memo map[int]int = map[int]int{}


func main() {
	fmt.Println(fibDynamic(25))
}

// Benchmark_fibonacci-8   	    7515	    154866 ns/op	       0 B/op	       0 allocs/op
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Benchmark_memoization-8   	229151509	         5.057 ns/op	       0 B/op	       0 allocs/op (If it includes fibonacci(recursive))
// Benchmark_memoization-8         102108819               11.39 ns/op            0 B/op          0 allocs/op (If it includes memoization)
// Memoization: Top Down
func memoization(n int)int{
	_, ok := memo[n]
	
	if ok {

        return memo[n]

    }
	if n <= 2 {
		return 1
	}
	memo[n]= memoization(n-1) + memoization(n-2)
	return memo[n]
}

// Benchmark_tabulation-8   	11276450	        98.55 ns/op	     224 B/op	       1 allocs/op
// Tabulation: Bottom Up
func tabulation(n int)int{
	var result []int = make([]int, n+2)

	result[0] = 0
    result[1] = 1


    for i := 0; i <= n-1; i++ {

        result[i+1] += result[i]

        result[i+2] += result[i]

    }


    return result[n]
}

// Benchmark_fibDynamic-8           9664682               123.0 ns/op           496 B/op          5 allocs/op
func fibDynamic(n int) int {
    a := []int{1}
    if n == 1 {
        return 1
    }
    a = append(a, 1)
    if n == 2 {
        return 2
    }
    for i := 2; i < n; i++ {
        a = append(a, a[i-2]+a[i-1])
    }
    return a[len(a)-1]
}