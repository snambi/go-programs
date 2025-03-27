package greetings

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type rectangle struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perim() float64 {
	return 4 * s.side
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perim() float64 {
	return (2 * r.length) + (2 * r.width)
}

func Measure(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perim())
}

func DetectCircle(s shape) {
	c, ok := s.(circle)
	if ok {
		fmt.Println("Circle detected")
		fmt.Println(c.radius)
	}
	r, ok := s.(rectangle)
	if ok {
		fmt.Println("Rectangle detected")
		fmt.Println(r.length)
		fmt.Println(r.width)
	}
}
