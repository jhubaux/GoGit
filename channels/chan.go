package channels

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func ComputeFibonacci(ns []int, ch int) {
	for _, n := range ns {
		go func(nb int) {
			ch <- Fibonacci(nb)
		}(n)
	}
}

func square(n int) int {
	return n * n
}

func cubic(n int) int {
	return square(n) * n
}

// what is the name in english for ^4 ?
func power4(n int) {
	return cubic(n) * n
}

func FanOut() (chan int,chan int,chan int) {
	cub := make(chan int)
	p4 := make(chan int)

	// square channel
	sq := make(chan int)
	go func() {
		for n := range sq {

		}
	}

	// main master function
	go func() {
		for n := range master {
			
		}
	}
}
