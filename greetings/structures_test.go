package greetings

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	x := Animal{"dog", 2}
	y := Animal{"cat", 3}

	fmt.Println("Before swap: ", x, y)
	Swap(&x, &y)

	fmt.Println("After swap: ", x, y)
	//fmt.Println("After swap: ", a, b)
}
