package main

import (
	"fmt"
)

const (
	greeting = "Hello, I'm a duplicator programm"
	message  = "Enter a number"
)

func main() {
	fmt.Println(greeting)
	fmt.Println(message)

	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
