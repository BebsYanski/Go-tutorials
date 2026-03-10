package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Min(4, 27))
	fmt.Println(math.Max(4, 27))
	fmt.Println(math.Ceil(4.2625))

	var x = "12345"
	y, err := strconv.Atoi(x)            // string to integer
	z, err := strconv.ParseInt(x, 10, 0) // string to integer

	fmt.Println(y, err, z)

}
