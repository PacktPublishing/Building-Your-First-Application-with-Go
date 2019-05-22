package main

import (
	"fmt"
	//"strings"
)

func main() {
	var x int
	fmt.Scanf("%d", &x)

	switch {
	case x >= 30:
		fmt.Println("Old")
		break
	case x >= 18 && x < 30:
		fmt.Println("Young")
		break
	case x < 18:
		fmt.Println("Kid")
		fallthrough
	default:		
		fmt.Println("Sorry, you are too young")
	}

	/*
	var direction string
	fmt.Scanf("%s", &direction)

	switch direction = strings.ToLower(direction); direction {
	case "n":
		fmt.Println("Go north!")
	case "s":
		fmt.Println("Go south!")
	case "w":
		fmt.Println("Go west!")
	case "e":
		fmt.Println("Go east!")
	default:
	}
	*/
}
