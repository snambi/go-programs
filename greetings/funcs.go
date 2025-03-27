package greetings

import "fmt"

func Fact(n int) int {

	if n == 0 {
		return 1
	} else if n < 0 {
		return -1
	}

	return n * Fact(n-1)
}

func FuncFunc(n int) int {

	mult := func(x int, y int) int {
		return x * y
	}

	x := mult(3, 4)

	fmt.Println("x = ", x)

	return x
}
