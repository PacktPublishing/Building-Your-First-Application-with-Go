package main

import (
	"fmt"
	"time"
)

func FanIn(ch1, ch2 <-chan int) chan int {
	out := make(chan int)
	go func() {
		for {
			select {
			case num := <- ch1:
				out <- num
			case num := <- ch2:
				out <- num
			}
		}
	}()
	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	manhole := FanIn(ch1, ch2)
	for {
		select {
			case num := <- manhole:
				fmt.Println(num)
			case <- time.After(time.Second * 2):
				fmt.Println("timeout")
				return
		}
	}
}
