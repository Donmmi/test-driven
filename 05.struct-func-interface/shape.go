package main

import "math"

type Perimeter interface {
	Perimeter() float64
}

type Area interface {
	Area() float64
}

type Shape interface {
	Perimeter
	Area
}

type Rectangle struct {
	height float64
	weight float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.weight + r.height) * 2
}

func (r Rectangle) Area() float64 {
	return r.weight * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {

}
