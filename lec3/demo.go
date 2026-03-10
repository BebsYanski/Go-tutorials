package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var x = 7
	var y = 1000

	var z = float64(y) / float64(x)
	var p = fmt.Sprint(y) + "Hey"

	fmt.Println(z, p)
}
