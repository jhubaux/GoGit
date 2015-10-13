package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Test if the fibonacci function is correct
func TestFibonacciSingle(t *testing.T) {
	val := Fibonacci(4)
	if val != 3 {
		t.Error(fmt.Sprintf("Fib(4) = 3 but had %d", val))
	}
}

// Given a waitgroup obj, you must implement function
// that takes a list of number and use goroutines to compute
// the fibonacci out of them
func TestWaitGroupFibonacci(t *testing.T) {
	//  Numbers to compute
	numbers := []int{4, 8, 16, 32, 20}
	var fibs []int
	// the WaitGroup is kinda like semaphore
	var wg sync.WaitGroup
	// you add the numbers of jobs you want to wait
	wg.Add(len(numbers))

	// Just here to timeout the whole thing.
	// More in channels/*
	done := make(chan bool)
	go func() {
		// you must impl this function so each time a number is computed
		// you call wg.Done() on it. If len(numbers) routines calls wg.Done()
		// the execution flow will return
		// WARNING : it should not block !!
		fibs = ComputeFibonacci(numbers, &wg)
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// verify that you have done the right computation.
		for i, n := range fibs {
			if fib := Fibonacci(numbers[i]); fib != n {
				t.Error(fmt.Sprintf("Numbers[%d] = %d but it should be equal to Fib(%d) = %d\n", i, n, numbers[i], fib))
			}
		}
	case <-time.After(time.Second * 2):
		t.Error("Fibonaccis numbers should have been computed by now ... !!")
	}
}
