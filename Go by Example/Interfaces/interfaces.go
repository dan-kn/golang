package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct{ radius float64 }

func (r Rectangle) area() float64      { return r.width * r.height }
func (r Rectangle) perimeter() float64 { return 2*r.width + 2*r.height }
func (c Circle) area() float64         { return math.Pi * c.radius * c.radius }
func (c Circle) perimeter() float64    { return 2 * math.Pi * c.radius }

func measure(g Geometry) {
	fmt.Println("area:", g.area())
	fmt.Println("perimeter:", g.perimeter())
}

func detectRectangle(g Geometry) {
	if r, ok := g.(Rectangle); ok {
		fmt.Println("it's a rectangle!")
		fmt.Println("width:", r.width)
		fmt.Println("height:", r.height)
	}
}

func detectCircle(g Geometry) {
	if c, ok := g.(Circle); ok {
		fmt.Println("it's a circle!")
		fmt.Println("radius:", c.radius)
	}
}

func main() {
	fmt.Println("...")

	rectangle := Rectangle{width: 3, height: 4}
	circle := Circle{radius: 5}

	measure(rectangle)
	measure(circle)

	detectRectangle(rectangle)
	detectCircle(circle)
}
