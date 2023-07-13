package main

import (
	"fmt"
	"math"
)

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

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}
func (s Square) Perimeter() float32 {
	return 4 * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	square := Square{5}
	circle := Circle{6}
	printArea(square)
	printArea(circle)
	printInterface(square)
	printInterface(circle)
	// printInterface(true)
	// printInterface(123)
	// printInterface("Wow")

}

func printArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	// fmt.Printf("%+v\n", i)
	switch t := i.(type) {
	case int:
		fmt.Println("int ", t)
	case bool:
		fmt.Println("bool ", t)
	default:
		fmt.Println("unknown type", t)
	}

	str, ok := i.(string)
	if !ok {
		fmt.Println("interface not string")
	} else {
		fmt.Println("len of i as string is ", len(str))
	}

}
