package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{}
type square struct{}

func main() {
	t := triangle{}
	s := square{}

	printArea(t)
	printArea(s)

}

func printArea(s shape) {
	fmt.Println(s.getArea())

}

func (t triangle) getArea() float64 {
	base := 4.0
	height := 3.0

	return base * height / 2
}

func (s square) getArea() float64 {
	length := 8.0
	width := 4.0

	return length * width
}
