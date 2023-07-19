package main

import "fmt"

type Value interface {
}

type User struct {
	emal string
	name string
}

type Cash struct {
	cash map[string]interface{}
}

func (c *Cash) New() {
	c.cash = make(map[string]interface{})
}
func (c *Cash) Set(key string, i interface{}) {
	c.cash[key] = i
}
func (c *Cash) Get(key string) (interface{}, bool) {
	i, ok := c.cash[key]
	return i, ok
}

func main() {

	var inCash Cash

	//j := int(10)
	// fmt.Println(j)
	u := User{"13@gmail.com", "Urah"}
	inCash.New()
	inCash.Set("user11", 45)
	i, ok := inCash.Get("user11")
	fmt.Println(i, ok)
	inCash.Set("user12", u)
	y, ok := inCash.Get("user12")
	fmt.Println(y, ok)

}

func makeMapV[V Value]() map[string]V {
	return make(map[string]V)
}

func foo(v Value) {
	//
}
