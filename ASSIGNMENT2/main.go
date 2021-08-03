package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
 t1 := triangle{height:1.5, base:1.2}
 s1 := square{sideLength: 2.1}

 printArea(t1)
 printArea(s1)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base) * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}