package main

import "fmt"

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func (u *User) printUserInfo() {
	fmt.Println(u.name, u.sex, u.age, u.height, u.weight)
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    Age(age),
		sex:    sex,
		weight: weight,
		height: height,
	}
}

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

func main() {
	user1 := User{"Joe", 32, "Male", 93, 182}
	// fmt.Printf("%+v\n", user1)
	user1.printUserInfo()
	fmt.Println(user1.age.isAdult())
	user2 := User{"Bob", 44, "Male", 80, 190}
	//fmt.Printf("%+v\n", user2)
	user2.printUserInfo()
}
