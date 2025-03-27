package greetings

import "fmt"

type base struct {
	x int
}

type dereived struct {
	base
	y int
	x base
}

func Hello2() string {

	x := dereived{
		base: base{
			x: 1,
		},
		y: 2,
		x: base{
			x: 3,
		},
	}

	fmt.Println("x = ", x)
	return "Hello, world."
}
