package greetings

import "fmt"

func Arrays_f() {
	var a [5]int

	fmt.Println("a = ", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var f = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", f)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", c)

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s))

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
}
