package greetings

func byval(x int) int {
	return x * 2
}

func byref(x *int) {
	*x = *x * 2
}
