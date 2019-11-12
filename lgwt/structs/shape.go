package structs

import "math"

type Rectangle struct {
	Width	float64
	Height	float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius	float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base		float64
	Height		float64
}

func (t Triangle) Area() float64 {
	return t.Base * ( 0.5 * t.Height)
}

type Shape interface {
	Area()	float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}