package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Gladys"))
}
