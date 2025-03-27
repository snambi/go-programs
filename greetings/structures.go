package greetings

type List[T any] struct {
	head, tail *Element[T]
}

type Element[T any] struct {
	next  *Element[T]
	value any
}

type Animal struct {
	Name string
	Age  int
}

func (A Animal) makeSound() string {
	return "Animal sound"
}

func Swap(x *Animal, y *Animal) {

	t := x.Name
	x.Name = y.Name
	y.Name = t
}
