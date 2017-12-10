package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
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

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("My name is:", p.Name)
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

	fmt.Println(totalArea(&c, &r))

	a := new(Android)
	a.Model = "Simple"
	a.Person.Name = "Robot"
	fmt.Println(a)
	a.Talk()
}
