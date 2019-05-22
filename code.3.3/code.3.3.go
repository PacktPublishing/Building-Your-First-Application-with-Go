package main

import "fmt"

func main() {
	var i, x int

	for {
		i++
		if i == 10 {
			break
		}
		if i % 2 != 0 {
			x += i
		} else {
			continue
		}
	}

	fmt.Println(x)
}
