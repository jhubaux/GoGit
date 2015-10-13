package goroutines

// I'll let you figure this out ...;)
func Fibonacci(n int) int {

}

// Fill that function so that for each numbers in ns:
//  - it launches a goroutines to compute the fibonacci number out of it
//  - when finished it call wg.Done()
// This function is non blocking, so you must return before the end of it !
// The numbers will be verified using the returned slice
// It's not strictly a good concurrent programming pattern but it's without
// channels, that's why. If that is boring you, go to the channel exercice !!
func ComputeFibonacci(ns []int, wg *sync.WaitGroup) []int {

}
