package main

import (
	"testing"
	"math"
	"fmt"
)

func TestPerimeter(t *testing.T) {
	perimeters := []struct{
		name  string
		shape Shape
		want  float64
	} {
		{
			"rectangle",
			Rectangle{10, 20},
			60.0,
		},
		{
			"circle",
			Circle{10},
			20 * math.Pi,
		},
	}

	for _, p := range perimeters {
		t.Run(fmt.Sprintf(p.name), func(t *testing.T) {
			assertPerimeter(t, p.shape, p.want)
		})
	}
}

func TestArea(t *testing.T) {
	areas := []struct{
		name  string
		shape Shape
		want  float64
	} {
		{
			"rectangle",
			Rectangle{10, 20},
			200.0,
		},
		{
			"circle",
			Circle{10},
			100 * math.Pi,
		},
	}
	for _, a := range areas {
		t.Run(fmt.Sprintf(a.name), func(t *testing.T) {
			assertArea(t, a.shape, a.want)
		})
	}
}

func assertPerimeter(t *testing.T, s Shape, want float64) {
	t.Helper()
	if s.Perimeter() != want {
		t.Errorf("got:[%f], expected:[%f]", s.Perimeter(), want)
	}
}

func assertArea(t *testing.T, s Shape, want float64) {
	t.Helper()
	if s.Area() != want {
		t.Errorf("got:[%f], expected:[%f]", s.Area(), want)
	}
}