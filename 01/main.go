package main

import (
	"fmt"
	"reflect"
)

func main() {

	const mess string = "III"
	var mess1 = 12

	message := ""
	message = "Hello world!"

	mess2 := []byte("asd")

	var a byte = 62
	fmt.Printf("%c\n", a)

	fmt.Println(reflect.TypeOf(mess1))
	fmt.Println(message)
	fmt.Println(mess2)

	a1, b, c := 1, 2, "str"
	a1, b = b, a1
	fmt.Println(a1, b, c)

}
