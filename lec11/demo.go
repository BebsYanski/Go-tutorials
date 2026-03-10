package main

import "fmt"

func main() {
	fmt.Println("Maps in GO")
	var mp map[string]int = map[string]int{"a": 1}
	np := map[string]int{"b": 4, "c": 2}
	emptyMap := map[string]int{} //empty
	mpMake := make(map[string]int)
	sliceMap := map[string][]int{} // key = string, value = slice of int
	slice2Map := map[string][]int{"a": {1, 2, 3}, "b": {2, 4, 6}}
	sliceMap["a"] = []int{1, 2, 3} // add or append to map
	// delete(slice2Map, "b")         // deleting from the map mutates the map in place

	for t, y := range slice2Map {

		slice2Map[t] = append(y, 50)
	}
	fmt.Println(mp, np, emptyMap)
	fmt.Println(mpMake, sliceMap, slice2Map)

	// Check if key exists in map
	value, ok := np["b"]
	fmt.Println(value, ok)

	// example
	newMap := map[uint]uint{}

	n := uint(100)
	// count := 1
	// for number := uint(1); number <= n; number++ {
	// 	if number%5 == 0 {
	// 		newMap[uint(count)] = uint(number)
	// 		count++
	// 	}
	// }

	// fmt.Println(newMap)

	// example 2
	for number := uint(1); number <= n; number++ {
		for d := uint(1); d <= 5; d++ {
			if number%d == 0 {
				newMap[d]++
			}
		}
	}

	fmt.Println(newMap)

}
