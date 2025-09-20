package main

import (
	"fmt"
	"math"
)

const str string = "constant"

func main() {
	fmt.Println(str)

	const n = 100000

	const val = 3e20 / n
	fmt.Println(val)

	fmt.Println(int64(val))

	fmt.Println(math.Cos(n))
	fmt.Println(math.Sin(n))
}
