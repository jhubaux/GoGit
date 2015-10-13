package channels

import (
	"fmt"
	"sync"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func ComputeFibonacci(ns []int, ch chan int) {
	var wg sync.WaitGroup
	for _, n := range ns {
		// signal we are launching a job
		wg.Add(1)
		go func(nb int) {
			ch <- Fibonacci(nb)
			// signal we are done
			wg.Done()
		}(n)
	}

	// wait for the end TO CLOSE THE CHANNEL
	// This is very important otherwise the test will waits for values never
	// coming
	// Since this function is supposed to be non blocking
	// we have to do it in a goroutine
	go func() {
		wg.Wait()
		close(ch)
	}()
}

// Takes an input channel and return a new channel
// with elements squared
func square(ch chan int) chan int {
	sq := make(chan int)
	go func() {
		for n := range ch {
			sq <- n * n
		}
		fmt.Printf("Closing channel square")
		// When the upper channel is closed, close this one
		close(sq)
	}()
	return sq
}

func cubic(ch chan int) chan int {
	cub := make(chan int)
	go func() {
		for n := range ch {
			cub <- n * n * n
		}
		close(cub)
	}()
	return cub
}

func FanOut(numbers []int) (chan int, chan int, chan int) {
	// Channels used to dispatch numbers
	sq_ := make(chan int)
	cub_ := make(chan int)
	p4_ := make(chan int)

	// Returned channels
	sq := square(sq_)
	cub := cubic(cub_)
	p4 := square(square(p4_))

	// Dispatch numbers to each channels
	go func() {
		for _, n := range numbers {
			sq_ <- n
			cub_ <- n
			p4_ <- n
		}
		// Then closes the channels
		close(sq_)
		close(cub_)
		close(p4_)
	}()

	return sq, cub, p4
}
