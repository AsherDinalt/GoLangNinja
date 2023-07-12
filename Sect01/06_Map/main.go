package main

import "fmt"

func main() {
	users := map[string]int{ // [key]value
		"John": 15,
		"Bob":  23,
		"Nick": 32,
	}
	fmt.Println(users["John"])
	age, ok := users["Mike"]

	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Not found")
	}

	users["Serge"] = 21
	delete(users, "Bob")

	for key, value := range users {
		fmt.Println(key, value)
	}

	var friends map[string]int
	friends = make(map[string]int)
	friends["Victor"] = 23

}
