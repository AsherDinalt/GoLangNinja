package shape

import "math"

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func NewSquare(length float32) Square {
	return Square{
		sideLength: length,
	}
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}
func (s Square) Perimeter() float32 {
	return 4 * s.sideLength
}

type Circle struct {
	radius float32
}

func NewCircle(radius float32) Circle {
	return Circle{
		radius: radius,
	}
}
func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}
