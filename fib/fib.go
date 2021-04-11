package fib

import (
	// "fmt"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibMemo(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}