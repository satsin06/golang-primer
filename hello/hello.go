package main

import (
	"fmt"

	"example.com/greetings"
)

// import "rsc.io/quote"

func main() {
	// fmt.Println("Hello, World!")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}