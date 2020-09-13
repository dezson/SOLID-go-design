package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type calculator struct{}

func (c calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func main() {
	c := circle{radius: 5}
	s := square{length: 6}

	cal := calculator{}
	fmt.Printf("areaSum = %f\n", cal.areaSum(c, s))
}
