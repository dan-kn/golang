package main

import "fmt"

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()

	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}

func intSeq() func() int {
	j := 0
	return func() int {
		j++
		return j
	}
}
