package main

import "fmt"

type bot interface {
	getGreeting() string
}

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

type rectangle struct {
	height float64
	width  float64
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there"

}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hola"

}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func printArea(s shape) {
	fmt.Println(s.area())
}
