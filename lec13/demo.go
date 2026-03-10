package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	p1 := Person{
		name: "Yan",
		age:  23,
	}

	p1.f = func(s string) string {
		return s + "s"
	}

	fmt.Println(p1)
	p1.name = "Tim"
	fmt.Println(p1.name)

	fmt.Print(p1.f("Good"))
	value := p1.GetName()

	fmt.Println(value)

	p1.SetName("Joey")
	fmt.Println(p1)

	//

	p2 := Person{
		age:      23,
		name:     "Yan",
		favSport: Sport{"Soccer", "D"},
	}

	fmt.Println(p2)

	//
	s1 := Student{
		name:     "Yan",
		age:      23,
		favSport: []Sport{{"Soccer", "D"}, {"Basket ball", "C"}},
	}

	fmt.Println(s1)

}

type Sport struct {
	name     string
	position string
}

type Person struct {
	name     string
	age      uint
	f        func(string) string
	favSport Sport
}

type Student struct {
	name     string
	age      uint
	favSport []Sport
}

// Method for a struct
func (p Person) GetName() string {
	return p.name
}

func (p Person) SetName(name string) {
	p.name = name
	fmt.Println(p)
}
