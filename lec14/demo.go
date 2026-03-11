package main

import "fmt"

func main() {
	fmt.Println("Interfaces in Go!")
	var s Shape
	s = Rectangle{width: 5, height: 10}
	fmt.Printf("Area of Rectangle: %.2f\n", s.Area())

	var p []Shape
	p = append(p, Rectangle{width: 5, height: 10})
	p = append(p, Circle{radius: 7})
	p = append(p, Square{side: 4})
	p = append(p, Triangle{base: 6, height: 8})

	for _, shape := range p {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}

	var i interface{}
	i = "Hello, World!"
	fmt.Println(i)

	var j []Shape = []Shape{
		Rectangle{width: 5, height: 10},
		Circle{radius: 7},
		Square{side: 4},
		Triangle{base: 6, height: 8},
	}
	for _, shape := range j {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}

}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
