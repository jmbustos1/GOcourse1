package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	tr := triangle{}
	tr.base = 5
	tr.height = 5

	sq := square{}
	sq.sideLength = 5

	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(tr)
	printArea(sq)

	printArea(t)
	printArea(s)

}
