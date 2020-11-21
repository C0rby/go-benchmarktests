package main

import "fmt"

func FibIterative(n int) int64 {
	var a, b int64
	a, b = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func FibRecursive(n int64) int64 {
	if n <= 1 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

/*
def fib(n, computed = {0: 0, 1: 1}):
...     if n not in computed:
...         computed[n] = fib(n-1, computed) + fib(n-2, computed)
...     return computed[n]
*/
func FibMemoization(n int) int64 {
	computed := map[int]int64{
		0: 0,
		1: 1,
	}
	var fib func(int, map[int]int64) int64
	fib = func(n int, computed map[int]int64) int64 {
		if _, ok := computed[n]; !ok {
			computed[n] = fib(n-1, computed) + fib(n-2, computed)
		}
		return computed[n]
	}

	return fib(n, computed)
}

func main() {
	fmt.Println(FibIterative(30))
	fmt.Println(FibRecursive(30))
	fmt.Println(FibMemoization(30))
}
