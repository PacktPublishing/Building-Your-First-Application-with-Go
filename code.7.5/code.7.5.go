package main

import "fmt"

func GoDivide(x int, y int) int {
	return x / y
}

func main() {
	defer func() {
		fixed := recover()
		fmt.Println(fixed)
	}()
	fmt.Println(GoDivide(1, 0))
}
