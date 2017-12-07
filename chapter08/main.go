package main

import (
	"fmt"
)

func main() {
	x1 := 1
	zeroV(x1)
	fmt.Println(x1)

	x2 := new(int)
	*x2 = 1
	zeroP(x2)
	fmt.Println(*x2)
}

func zeroV(x int) {
	x = 0
}

func zeroP(xPtr *int) {
	*xPtr = 0
}
