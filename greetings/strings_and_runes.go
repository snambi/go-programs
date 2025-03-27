package greetings

import "fmt"

func StringsAndRunes() {
	s := "hello"
	fmt.Println("s = ", s)
	fmt.Println("s[0] = ", s[0])
	fmt.Println("s[1:3] = ", s[1:3])
	fmt.Println("len(s) = ", len(s))
	fmt.Println("s + world = ", s+" world")

	for i, c := range s {
		fmt.Println("i = ", i, "c = ", c)
	}

	r := 'a'
	fmt.Println("r = ", r)
	fmt.Println("r = ", string(r))

	t := "こんにちは"
	fmt.Printf("%q\n", t)
	fmt.Printf("%+q\n", t)
	fmt.Printf("% x\n", t)
}
