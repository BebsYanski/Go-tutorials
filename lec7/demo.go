package main

import "fmt"

func main() {
	fmt.Println("Hello")

	for x := 0; x <= 10; x++ {
		fmt.Println(x)
	}

	// There is no while loop in Go. We use the FOR loop for that
	a := 10

	for a < 10 {
		fmt.Println(a)
		a++
	}
}
