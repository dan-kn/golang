package main

import (
	"fmt"
	"slices"
)

func main() {
	var a [3]int
	fmt.Println("emp:", a)

	a[1] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[1])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [3][2]int
	for i := range 3 {
		for j := range 2 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println("2d: ", twoD)
}
