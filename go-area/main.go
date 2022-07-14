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
	side float64
}

func main() {
	triangle := triangle{
		base:   3,
		height: 5,
	}

	square := square{
		side: 3,
	}
	printArea(square)
	printArea(triangle)
}

func (t triangle) getArea() float64 {
	area := .5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.side * s.side
	return area
}

func printArea(s shape) float64 {
	fmt.Println(s.getArea())
	return s.getArea()
}
