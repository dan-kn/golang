package main

import "fmt"

func main() {
	fmt.Println("something goes here...")

	xVal := 32
	yVal := 10

	fmt.Println(addVals(xVal, yVal))
	fmt.Println(subVals(xVal, yVal))
	fmt.Println(multVals(xVal, yVal))
	fmt.Println(isEven(xVal))

	zVal := 128

	fmt.Println(addThreeVals(xVal, yVal, zVal))
}

func addVals(xVal int, yVal int) int                { return xVal + yVal }
func subVals(xVal int, yVal int) int                { return xVal - yVal }
func multVals(xVal int, yVal int) int               { return xVal * yVal }
func isEven(xVal int) bool                          { return xVal%2 == 0 }
func addThreeVals(xVal int, yVal int, zVal int) int { return xVal + yVal + zVal }
