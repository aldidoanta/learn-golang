package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		base:   2,
		height: 6,
	}

	s := square{sideLength: 2}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("the area is ", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
