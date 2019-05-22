// Main package - the heart of your application
package main

// Import fmt package to use input/output functions
import "fmt"

/*
  Main function - an entry point for the application.
  This function's body contains all needed code.
  It receives input, processes data, and produces output.
*/
func main() {
	var radius float32
	// Receiving input
	fmt.Scanf("%f", &radius)

	var length float32
	// Calculating the result
	length = 2 * 3.14 * radius

	// Making output
	fmt.Println("Length:", length)

	// TODO: Make a cool application!
}
