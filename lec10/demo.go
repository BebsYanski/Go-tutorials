package main

import "fmt"

func main() {

	fmt.Println("Slices")
	// Much more flexible than arrays. They can be resized i.e grow dynamically
	// It is a view of the array. You can increase, decrease array capacity
	// A slice has a pointer to the underlying array.
	// Made up of: Pointer (slice[0]) to it's first element, Length (len) of the slice, and capacity of the slice (cap(slice)) which how much the slice can store
	arr := [5]int{1, 2, 3, 4, 5}

	s1 := arr[:]
	s2 := arr[1:3]
	s3 := arr[2:]
	// change to a slice affects the main array
	s1[2] = 100

	fmt.Println(arr, s1, s2, s3)
	fmt.Printf("%T %T %T %T\n", arr, s1, s2, s3)

	// Pointer
	// length
	// capacity
	fmt.Println(s2, len(s2), cap(s2))

	s5 := [2]string{"Hello", "World"} // This is an array
	s4 := []string{"Hello", "World"}  // This is a slice, not an array
	fmt.Println(s4, s5)
	fmt.Printf("%T, %T\n", s4, s5)

	for i := 0; i < 10; i++ {
		fmt.Println(s4, len(s4), cap(s4))
		s4 = append(s4, "Yan"+fmt.Sprint(i))
	}

	// other ways to make slices
	s6 := make([]int, 10) // creates an integer slice of size 10
	fmt.Println(s6)
	s7 := []string{"Hello", "world", "hi"}
	for _, y := range s7 {
		fmt.Println(y)
	}

	// Notice that unlike an array, the original slice can be mutated within a function
	s8 := s5[:]
	test(s8)
	fmt.Println(s8, s5)
}

func test(arr []string) {
	arr[0] = "Change this"
}
