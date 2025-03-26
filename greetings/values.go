package greetings

import (
	"fmt"
	"math"
	"time"
)

func Values_f() {
	fmt.Println(" a + b =", ("a" + "b"))
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

const s string = "constant"

func Const_f() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func Switch_f() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	j := time.Now().Weekday()
	fmt.Println("J ", j)

	switch j {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	WhatAmI := func(t interface{}) {
		fmt.Println("I am a type= ", t)
	}

	WhatAmI(true)
	WhatAmI(3)
	WhatAmI("hey")
}
