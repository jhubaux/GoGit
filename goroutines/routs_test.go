package goroutines

import (
	"sync"
)

// Test if the fibonacci function is correct
func TestFibonacciSingle(t *testing.T) {
	val := Fibonacci(4)
	if val != 24 {
		t.Error(fmt.Sprintf("Fib(4) = 24 but had %d", val))
	}
}

// Given a waitgroup obj, you must implement function
// that takes a list of number and use goroutines to compute
// the fibonacci out of them
func TestWaitGroupFibonacci(t *testing.T) {
	//  Numbers to compute
	numbers := []int{4, 8, 16, 32, 56}
	fibs := make([]int, len(numbers))
	// the WaitGroup is kinda like semaphore
	var wg sync.WaitGroup
	// you add the numbers of jobs you want to wait
	wg.Add(len(numbers))

	// you must impl this function so each time a number is computed
	// you call wg.Done() on it. If len(numbers) routines calls wg.Done()
	// the execution flow will return
	ComputeFibonacci(numbers, fibs, &wg)

	// Just here to timeout the whole thing.
	// More in channels/*
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		fmt.Printf("Fibs : %v\n", fibs)
	case <-time.After(time.Second * 3):
		t.Error("Fibonaccis numbers should have been computed by now ... !!")
	}
}
