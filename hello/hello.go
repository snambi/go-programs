package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	message, err := greetings.Hello("xx")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println(message)

	// names := []string{"Gladys", "Samantha", "Darrin"}
	// msgs, err := greetings.Hellos(names)

	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// fmt.Println(msgs)
	//greetings.Values_f()
	//greetings.Const_f()
	//greetings.Switch_f()

	greetings.Arrays_f()
}
