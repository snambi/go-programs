package greetings

import "fmt"

func Hello(name string) string {
	c := 10
	fmt.Println(c)
	message := fmt.Sprintf("Hello, %v welcome!", name)
	return message
}
