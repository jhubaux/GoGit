package goroutines

import (
	"sync"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func ComputeFibonacci(ns []int, wait *sync.WaitGroup) []int {
	fibs := make([]int, len(ns))
	for i, n := range ns {
		// Close functions => i change over time so pass a copy of the index
		go func(j, nb int) {
			fibs[j] = Fibonacci(nb)
			wait.Done()
		}(i, n)
	}
	return fibs
}
