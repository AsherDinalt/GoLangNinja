package main

import (
	"errors"
	"fmt"
	// "reflect"
)

func main() {

	var message string
	message = "Hi everybody"

	printMessage(message)

	message = sayHello("Max", 21)
	printMessage(message)
	mess, ent := enterClub(60)
	fmt.Println(mess, ent)

	mess, err := enterClubErr(15)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s! You're %d years old", name, age)
}

func enterClub(age int) (string, bool) {
	if age >= 18 && age < 45 {
		return "Come in", true
	} else if age >= 45 {
		return "Are you sure?", true
	} else {
		return "Fuck off, bustard", false
	}
}

func enterClubErr(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Come in", nil
	} else if age >= 45 {
		return "Are you sure?", nil
	} else {
		return "Fuck off, bustard", errors.New("You're too small")
	}
}
