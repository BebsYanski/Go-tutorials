package main

import "fmt"

func main() {
	fmt.Println("arrays")

	// arrays here have a fixed size
	var arr [4]int
	numbers := [5]int{2, 4, 6, 8, 10}
	nested := [2][2]int{{2, 3}, {4, 5}}
	nestednew := [...][2]int{{2, 3}, {4, 5}, {6, 7}} // ... to not specify, as such it is infered

	fmt.Println(arr, len(arr), numbers, nested, nestednew)

	// Array properties
	arr[0] = 23
	nested[0] = [2]int{10, 11}
	fmt.Println(arr, len(arr), numbers, nested, nestednew)

	// Looping through array
	for i, value := range nested {
		fmt.Println("For index:", i, "We have:", value)
		for _, y := range value {
			fmt.Println(y)

		}
	}

	test(nested)
	fmt.Println(nested)
}

func test(arr [2][2]int) {
	fmt.Println("within func - before mutation", arr)
	arr[0] = [2]int{100, 100}
	fmt.Println("within func - after mutation", arr)

}
