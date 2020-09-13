package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() {
	fmt.Printf("Circle area %f\n", math.Pi*c.radius*c.radius)
}

type square struct {
	length float64
}

func (s square) area() {
	fmt.Printf("Square area %f\n", s.length*s.length)
}

func main() {
	c := circle{radius: 5}
	c.area()

	s := square{length: 6}
	s.area()
}
