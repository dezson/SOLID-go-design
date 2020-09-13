package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	name() string
}

type printer struct{}

func (p printer) displayArea(s shape) {
	fmt.Printf("The area of %s = %f \n", s.name(), s.area())
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) name() string {
	return "Circle"
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) name() string {
	return "Square"
}

func main() {

	p := printer{}
	p.displayArea(circle{radius: 5})
	p.displayArea(square{length: 6})
}
