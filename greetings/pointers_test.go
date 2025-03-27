package greetings

import (
	"fmt"
	"testing"
)

func TestByval(t *testing.T) {
	result := byval(5)
	expected := 10
	if result != expected {
		t.Errorf("byval(5) = %d; want %d", result, expected)
	}
}

func TestByref(t *testing.T) {
	x := 5
	byref(&x)
	expected := 10
	if x != expected {
		t.Errorf("byref(&x) = %d; want %d", x, expected)
	}
}

func TestByValRef(t *testing.T) {
	x := 4
	fmt.Println("X val before byval = ", x)
	y := byval(x)
	fmt.Println("X val after byval = ", x)
	fmt.Println("Y val after byval = ", y)

	byref(&x)
	fmt.Println("X val after byref = ", x)
}
