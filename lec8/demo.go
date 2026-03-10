package main

import "fmt"

func main() {
	fmt.Println("iterating strings in Go")
	// Go stores string characters as bytes
	// a byte is uint8
	str := "Hello World"
	fmt.Printf("%T, %d\n", str[4], str[4])

	// To get it as string:
	fmt.Println(string(str[4]))

	for x := 0; x < len(str); x++ {
		// Type here is uint8 := byte
		fmt.Printf("\n%c\tType:%T", str[x], str[x])
		// fmt.Println(string(str[x]))

	}

	// Clearer, recommended syntax
	for i, char := range str {
		// Type here is int32 := rune
		fmt.Printf("\nCharacter at index %d: %c\tType: %T", i, char, char)
	}
}
