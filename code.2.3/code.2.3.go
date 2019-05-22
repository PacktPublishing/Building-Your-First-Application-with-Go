package main

import "fmt"

var (
	x int
	px *int
)

func main() {
	x = 15
	px = &x

	*px = 19
	fmt.Println(x)
}
