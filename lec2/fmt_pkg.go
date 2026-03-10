package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// printf format specifiers: %d - decimal numbers
	// %T - data type of a variable/value
	// %v - The value itself
	// %b - get binary equivalent of a decimal
	// %e - get exponential form of a large decimal place
	// %f - floats
	// %s - String

	var x = 10.35342
	var y = fmt.Sprintf("%10.2f%%", x) //returns whatever you give as a string

	o, p := fmt.Println(y)
	fmt.Println(o, p)
}
