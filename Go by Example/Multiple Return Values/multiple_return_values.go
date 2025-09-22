package main

import "fmt"

func main() {
	fmt.Println("something goes here...")

	xVal, yVal := 32, 16

	fmt.Println("---BEFORE SWAP---")
	fmt.Println("xVal =", xVal, "| yVal =", yVal)

	xVal, yVal = swapVals(xVal, yVal)

	fmt.Println("---AFTER SWAP---")
	fmt.Println("xVal =", xVal, "| yVal =", yVal)

	fmt.Println("--- ... ---")

	_, yVal = vals()
	fmt.Println("yVal =", yVal)
}

func swapVals(xVal int, yVal int) (int, int) { return yVal, xVal }
func vals() (int, int) { return 100, 200 }
