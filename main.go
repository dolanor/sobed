package main

import (
	"fmt"

	"github.com/dolanor/sobed/greet"
)

func main() {
	greetFn, err := dislodgeAndDLOpen(greet.LibFS, "libgreet.so", "greet")
	if err != nil {
		panic(err)
	}

	greeter := CGreeter{
		greetFn: greetFn,
	}
	greeting := greeter.Greet("World")
	fmt.Println(greeting)
}
