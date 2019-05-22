package main

import (
	"fmt"
)

func routine(in <-chan int, out chan<- bool) {
	for {
		i, ok := <- in
		if !ok  {
			break
		}
		sum := 0
		for j := 0; j < i; j++ {
			sum += j
		}
		fmt.Println("Sum =", sum)
	}
	out <- true
}

func main() {
	in := make(chan int, 5)
	out := make(chan bool)
	go routine(in, out)

	for i := 1000; i < 1020; i++ {
		fmt.Println("Sending:", i)
		in <- i
	}
	close(in)
	<- out
}
