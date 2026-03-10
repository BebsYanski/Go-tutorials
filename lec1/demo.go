package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// Data types in Go
	// Unsigned integers : uint8, 16, 32, 64
	// signed integers : int8, 16, 32, 64
	// floating point numbers : float32, 64
	// byte type: byte - for characters or int8
	// rune : for storing int32 values
	// boolean : bool
	// strings : string
	// nil : for undefined or none

	// Explicitly define when you do not want to assign it at once
	var myName string
	myName = "Yannick"
	var age uint16 = 23

	const x uint8 = 2

	guys := 2.3 // Implicitly definning
	y := uint(x)
	z := int64(y)

	fmt.Println(myName, age, guys, z)
	fmt.Printf("Type of %.2f:\t%T", guys, guys)
	// fmt.Printf("\t%T", guys)

}
