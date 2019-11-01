package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	length float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{length: 4}
	tr := triangle{height: 3, base: 4}

	printArea(sq)
	printArea(tr)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func printArea(s shape) {
	fmt.Println("Area =", s.getArea())
}
