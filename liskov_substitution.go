package main

import "fmt"

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	degree string
	salary float64
}

type student struct {
	human
	grades map[string]int8
}

type person interface {
	getName() string
}

type printer struct{}

func (pr printer) info(p person) {
	fmt.Println("Name ", p.getName())
}

func main() {
	h := human{name: "Adam"}
	s := student{
		human: human{name: "Eve"},
		grades: map[string]int8{
			"Math":    5,
			"Physics": 4,
		},
	}

	t := teacher{
		human:  human{name: "Mikkel"},
		degree: "CS",
		salary: 3000,
	}

	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
