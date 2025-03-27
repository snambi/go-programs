package greetings

func Range1() {
	x := []int{1, 2, 3, 4, 5}

	for i, v := range x {
		println("i = ", i, "v = ", v)
	}
}
