package main

import "fmt"

func main() {
	x := map[int]int{1: 10, 2: 20}
	fmt.Println("Value for 3:", x[3])

	_, ok := x[1]
	if ok {
		fmt.Println("Value for 1 exists")
	}
}
