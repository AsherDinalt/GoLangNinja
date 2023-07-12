package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(prediction("ss"))
	func() {
		fmt.Println("===========")
	}()

	inc1 := increment()
	inc2 := increment1(7)
	inc3 := increment2(10)
	fmt.Println(inc1())
	fmt.Println(inc2())

	fmt.Println(inc3(2))
	fmt.Println(inc3(2))
	fmt.Println(inc3(2))
	fmt.Println(inc3(3))

}

func increment() func() int {
	count := 0 // это в области видимости экземпляра
	return func() int {
		count++
		return count
	}
}

func increment1(start int) func() int {
	count := start // это в области видимости экземпляра
	return func() int {
		count++
		return count
	}
}

func increment2(start int) func(step int) int {
	count := start // это в области видимости экземпляра
	return func(step int) int {
		count += step
		return count
	}
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "Понедельник", nil
	case "вт":
		return "Вторник", nil
	case "ср":
		return "Среда", nil
	case "чт":
		return "Четверг", nil
	case "пт":
		return "Пятница", nil
	case "сб":
		return "Суббота", nil
	case "вс":
		return "Воскресенье", nil
	default:
		return "Не валидный день", errors.New("Not valid day")
	}
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
