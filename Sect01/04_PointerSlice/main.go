package main

import (
	"errors"
	"fmt"
)

var msg string

func init() { // выполняется ДО main()
	msg = " Line 1 "

}

func main() {
	// fmt.Println(prediction("ss"))
	// fmt.Println(msg)
	// printMsg(&msg)
	// fmt.Println(msg)

	// number := 5
	// var p *int
	// p = &number
	// fmt.Println(p, *p)

	//messages := [3]string{"1", "2", "3"}
	//messages := []string{"1", "2", "3"}
	//messages := make([]string)
	//fmt.Println(messages)
	//var msg []string
	// msg := make([]string, 5)
	// msg = append(msg, "1")
	// fmt.Println(msg)
	//printMsgS(msg)

	// var matrix = [][]int{}
	matrix := make([][]int, 10)
	counter := 0

	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			matrix[i][j] = counter
			counter++
		}
		//fmt.Println(matrix[i])
	}
	// mess := []string{
	// 	"m 01",
	// 	"m 02",
	// 	"m 03",
	// 	"m 04",
	// 	"m 05",
	// 	"m 06",
	// }
	// for i := range mess {
	// 	fmt.Println(mess[i])
	// }

	// for i, message := range mess {
	// 	fmt.Println(i, message)
	// }

	// for _, message := range mess {
	// 	fmt.Println(message)
	// }

	counter = 0

	for {
		counter++
		fmt.Println(counter)
		if counter > 100 {
			break
		}
	}

}

func printMsgS(msg []string) error {

	if len(msg) == 0 {
		return errors.New("empty")
	}

	msg[1] = "5" // в случау слайса происходит перезапись
	fmt.Println(msg)
	return nil
}

func printMsg(msg *string) {
	*msg += " and Line 2"
	fmt.Println(*msg)
}
