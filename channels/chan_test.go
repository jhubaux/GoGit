package channels

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Compute many fibonacci numbers and send them trough a channel
func TestFibonacciChannels(t *testing.T) {
	ch := make(chan int)

	numbers := []int{4, 8, 16, 32}
	// compute solutions
	solution := make(map[int]int)
	for _, n := range numbers {
		solution[Fibonacci(n)] = n
	}

	// just in case you wrote a blocking method .. which is bad !
	go func() {
		wait := make(chan bool)
		go func() {
			ComputeFibonacci(numbers, ch)
			wait <- true
		}()
		select {
		case <-wait:
		case <-time.After(time.Millisecond * 250):
			t.Error(fmt.Sprintf("Compute Fibonacci should not be a blocking operation"))
		}
	}()

	done := make(chan bool)
	// waiting for the inputs
	go func() {
		for fib := range ch {
			if _, ok := solution[fib]; !ok {
				t.Error(fmt.Sprintf("Received wrong number does not correspond to a fib number"))
			}
		}
		done <- true
	}()

	// timeout
	select {
	case <-done:
	case <-time.After(time.Second * 2):
		t.Error("Timeout the whole thing. What the heck are you doing ??")
	}
}

// Fan out method : one master channel -> dispatching the value to many children
// channels.
// Implement a method taking numbers in argument
// THe method should be returning 3 channels :
// first one returns n*n
// second one returns n*n*n
// third one returns n*n*n*n
// HENCE IT SHOULD NOT BLOCK !
// There are many ways to do this but this exercise illustrate what is the fan
// out pattern in Go. Your function should launch 3 goroutines each listening on
// their own channel. Your function should listen dispatch the numbers to the
// right channels and
// dispatch the value to the respective 3 children channels each time a new
// value is given.

// For more in-depth review of channels / pipelines :
// https://blog.golang.org/pipelines

func TestFanOut(t *testing.T) {
	numbers := []int{4, 8, 16, 32, 56}
	solution2 := make(map[int]int)
	solution3 := make(map[int]int)
	solution4 := make(map[int]int)
	for _, n := range numbers {
		solution2[n*n] = n
		solution3[n*n*n] = n
		solution4[n*n*n*n] = n
	}

	// gooo
	ch2, ch3, ch4 := FanOut(numbers)

	// waitgroup to wait every numbers to be generated and verified
	var wg sync.WaitGroup
	// verification function
	fn := func(ch chan int, sol map[int]int) {
		for n := range ch {
			if _, ok := sol[n]; !ok {
				t.Error(fmt.Sprintf("Received %d but that corresponds to nothing guy ", n))
			}
		}
		wg.Done()
	}
	go fn(ch2, solution2)
	go fn(ch3, solution3)
	go fn(ch4, solution4)

	// chan to use a timeout on the verification
	done := make(chan bool)
	// Wait the end  of the verification
	go func() {
		wg.Wait()
		// then signals the end
		done <- true
	}()
	select {
	case <-done:
	case <-time.After(time.Second * 2):
		t.Error("Too long !")
	}

}
