package structs

import "math"

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Length)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
