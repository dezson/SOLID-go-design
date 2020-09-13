package main

import (
	"errors"
	"fmt"
	"math"
)

type calculator struct{}

func (c calculator) areaSum(shapes ...interface{}) (float64, error) {
	var sum float64
	for _, shape := range shapes {
		switch shape.(type) {
		case circle:
			r := shape.(circle).radius
			sum += math.Pi * r * r
		case square:
			l := shape.(square).length
			sum += l * l
		default:
			return 0.0, errors.New("Invalid type")
		}
	}
	return sum, nil
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func main() {
	c := circle{radius: 5}
	s := square{length: 6}

	cal := calculator{}
	sum, err := cal.areaSum(c, s)
	if err != nil {
		fmt.Errorf("areaSum error %s", err)
	} else {
		fmt.Printf("areaSum = %f\n", sum)
	}
}
