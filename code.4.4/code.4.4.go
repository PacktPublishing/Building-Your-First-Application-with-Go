package main

import "fmt"

var m = make(map[int]string)

func init() {
	for key := 0; key < 10; key++ {
		m[key] = ""
	}
}

func main() {
	fmt.Println(m)
}
