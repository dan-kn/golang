package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) area() int     { return r.width * r.height }
func (r Rectangle) perimeter() int { return 2*r.width + 2*r.height }

func main() {
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Println("area:", rectangle.area())
	fmt.Println("perimeter:", rectangle.perimeter())

	rp := &rectangle
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
