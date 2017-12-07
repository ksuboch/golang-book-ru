package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Printf("%-3d ", i)
		if i%2 == 0 {
			fmt.Println("Divisible by 2")
		} else if i%3 == 0 {
			fmt.Println("Divisible by 3")
		} else if i%5 == 0 {
			fmt.Println("Divisible by 5")
		} else {
			fmt.Println("Unknown")
		}
		i = i + 1
	}

	for j := 1; j <= 10; j++ {
		fmt.Printf("%-3d ", j)
		if j%2 == 0 {
			fmt.Println(" odd")
		} else {
			fmt.Println(" even")
		}
	}

	for k := 1; k <= 10; k++ {
		fmt.Printf("%-3d ", k)
		switch k {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
		case 6:
			fmt.Println("six")
		case 7:
			fmt.Println("seven")
		case 8:
			fmt.Println("eight")
		case 9:
			fmt.Println("nine")
		case 10:
			fmt.Println("ten")
		}
	}
}
