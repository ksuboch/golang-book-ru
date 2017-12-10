# Задачи
### Какая разница между методом и функцией?
Метод - функция особого вида; для него необходимо указывать получателя. Получатель указывается между ключевым словом `func` и именем функции, после чего можно вызывать метод следующим образом:
```go
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

fmt.Println(c.area())
```
***
### В каких случаях могут пригодиться встроенные (скрытые) поля?
Они нужны для описания отношений вида "Объект А" является "Объектом Б"
***
### Добавьте новый метод perimeter в интерфейс Shape, который будет вычислять периметр фигуры. Имплементируйте этот метод для Circle и Rectangle.
```go
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rect struct {
	x1, y1, x2, y2 float64
}

func (r *Rect) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rect) perimeter() float64 {
	return 2 * (distance(r.x1, r.y1, r.x2, r.y1) + distance(r.x1, r.y1, r.x1, r.y2))
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c.perimeter())
	fmt.Println(c.area())

	r := Rect{0, 0, 10, 5}
	fmt.Println(r.x1, r.x2, r.y1, r.y2)
	fmt.Println(r.perimeter())
	fmt.Println(r.area())
}
```