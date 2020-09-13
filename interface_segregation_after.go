package main

import (
    "fmt"
    "math"
)

type square struct {
    length float64
}

func (s square) area() float64 {
    return s.length * s.length
}

type cube struct {
    square
}

func (c cube) area() float64 {
    return math.Pow(c.length, 2)
}

func (c cube) volume() float64 {
    return math.Pow(c.length, 3)
}

type shape interface {
    area() float64
}

type object interface {
    shape
    volume() float64
}

func areaSum(shapes ...shape) float64 {
    var sum float64
    for _, s := range shapes {
        sum += s.area()
    }
    return sum
}

func areaVolumeSum(shapes ...object) float64 {
    var sum float64
    for _, s := range shapes {
        sum += s.area() + s.volume()
    }
    return sum
}

func main() {
    s := square{length: 5}
    c := cube{s}
    fmt.Println(areaSum(s))
    fmt.Println(areaVolumeSum(c))
}
