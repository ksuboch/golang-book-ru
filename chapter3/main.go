package main

import "fmt"

func main() {
	fmt.Println("1 + 1     = ", 1+1)
	fmt.Println("1.0 + 1.0 = ", 1.0+1.0)

	fmt.Println("len(\"Hello World\") = ", len("Hello World"))
	fmt.Println("\"Hello World\"[1]   = ", "Hello World"[1])

	fmt.Println("true&&true  = ", true && true)
	fmt.Println("true&&false = ", true && false)
	fmt.Println("true||true  = ", true || true)
	fmt.Println("true||false = ", true || false)
	fmt.Println("!true       = ", !true)
}
