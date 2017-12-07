# Задачи
### Что делает следующий код?
```go
i := 10

if i > 10 {
    fmt.Println("Big")
} else {
    fmt.Println("Small")
}
```
Выводит в консоль сообщение `Small`, т.к. `i` не меньше 10
***
#### Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. (3, 6, 9, …).
```go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
```
***
### Напишите программу, которая выводит числа от 1 до 100. Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz», а для кратных как трём, так и пяти — «FizzBuzz».
***
```go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		res := ""

		if i%3 == 0 {
			res += "Fizz"
		}

		if i%5 == 0 {
			res += "Buzz"
		}

		if res != "" {
			fmt.Println(res)
		} else {
			fmt.Println(i)
		}
	}
}
```
