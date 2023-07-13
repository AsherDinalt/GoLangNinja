package main

import (
	"basic/shape"
	"fmt"
	"time"
)

func main() {
	// square := shape.NewSquare(5)
	// circle := shape.NewCircle(6)
	// printArea(square)
	// printArea(circle)
	// printInterface(square)
	// printInterface(circle)
	// printInterface(true)
	// printInterface(123)
	// printInterface("Wow")
	t := time.Now()
	fmt.Println(t)
}

func printArea(s shape.Shape) {
	fmt.Println(s.Area())
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
