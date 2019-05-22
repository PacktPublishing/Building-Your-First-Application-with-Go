package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("%d", &x)

	if message := "Hello, "; x >= 30 {
		fmt.Println(message, "Old")
	} else if x >= 18 && x < 30 {
		fmt.Println(message, "Young")
	} else {
		fmt.Println(message, "Sorry, you are too young")
	}
}
