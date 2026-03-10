package main

import "fmt"

func main() {
	fmt.Println("Functions")
	fmt.Println(add(3, 5))
	fmt.Println(square(3))
	fmt.Println(tripple(square, 5))

	// using anonymous functions
	value := tripple(func(x int) int {
		return 2 * x
	}, 3)

	fmt.Println(value)

	newFunc := func(x int) int {
		return square(x)
	}

	newVal := tripple(newFunc, 4)

	fmt.Println(newVal)

	// returning a function
	double := getFunc()
	fmt.Println(double(5))

	// returning a function with a parameter
	greet := getFunc2("Hello, ")
	fmt.Println(greet("World!"))

	// variadic functions
	fmt.Println(sum(1, 2, 3, 4, 5))

	sum := sum([]int{1, 2, 3, 4, 5}...) // unpacking a slice into variadic function
	fmt.Println(sum)

}

func add(a int, b int) (int, string) {
	return a + b, "hello"
}

func tripple(function func(int) int, num int) int {
	return function(num) * num
}

func square(a int) int {
	return a * a
}

func getFunc() func(int) int {
	return func(x int) int {
		return 2 * x
	}
}

func getFunc2(a string) func(string) string {
	return func(b string) string {
		return a + b
	}
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func sum2(nums []int) (total int, total2 int) {
	total = 0
	total2 = 0
	for _, num := range nums {
		total += num
	}
	return total, total2
}
