package main

import (
	"fmt"
)

func routine(i int, c chan int) {
	c <- i
}

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go routine(i, c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Received", i, <-c)
	}
}
