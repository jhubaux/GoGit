package channels

// Compute the fibonaccni number for each numbers and put the result into the
// chan. THis SHOULD NOT be a blocking function,i.e. it must return BEFORE the
// results are done
func ComputeFibonacci(ns []int, ch chan int) {

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

func FanOut(ns []int) (chan int, chan int, chan int)
