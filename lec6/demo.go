package main

import "fmt"

func main() {
	fmt.Println("Hello")

	a := 20

	switch a {
	case 1:
		fmt.Print("A = 1") // By default, Go breaks
	case 2:
		fmt.Println(a)
		fallthrough // allows you to not break, but extends to next case
	case 3:
		fmt.Println("Hehe")
	case 4, 5, 6:
		fmt.Println("Good")
	default:
		fmt.Println("Tada")

	}

	b := 20

	switch {
	case a < b:
		fmt.Println("cool")
	case b == 2:
		fmt.Println("True")
	case b == a, b == 34, b > a:
		fmt.Println("Going")
	default:
		fmt.Print("Done")
	}
}
