package main

import "fmt"

func main() {
	fmt.Println("hello, world!")

	sumVals(1, 2)
	sumVals(1, 2, 3)
	sumVals(1, 2, 3, 4)
	sumVals(1, 2, 3, 4, 5)

	nums := []int{1, 2, 3}
	sumVals(nums...)
}

func sumVals(vals ...int) {
	fmt.Println(vals, " ")

	sum := 0
	for _, val := range vals {
		sum += val
	}

	fmt.Println("sum:", sum)
}
