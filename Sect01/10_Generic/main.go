package main

import "fmt"



type Number interface {
	int64 | float64
}

func main() {
	a := []int64{1, 2, 3, 4}
	b := []float64{1.1, 2.2, 3.3}
	fmt.Println(sum(a))
	fmt.Println(sumN(b))
}

func sum[V int64 | float64](input []V) V {
	var res V
	res = 0
	for _, number := range input {
		res += number
	}
	return res

}

func sumN[V Number](input []V) V {
	var res V
	res = 0
	for _, number := range input {
		res += number
	}
	return res

}

func searchC[C comparable](els []C, el C) bool {
	for _, ell := range els {
		if ell == el {
			return true
		}
	}
	return false
}

func prnt[A any](dd A) {
	fmt.Println(dd)
}
