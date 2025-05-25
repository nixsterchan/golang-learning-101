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
	square := square{sideLength: 7}
	triangle := triangle{base: 4, height: 8}

	println("Area of the square is:")
	printArea(square)
	println("Area of the triangle is:")
	printArea(triangle)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
