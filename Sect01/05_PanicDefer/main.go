package main

import "fmt"

func main() {

	defer handlePanic()

	mess := []string{
		"m 01",
		"m 02",
		"m 03",
		"m 04",
		"m 05",
		"m 06"}
	mess[6] = "m 07"
	//panic("PANIC")
	fmt.Println(mess)

}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("After recover")
}
