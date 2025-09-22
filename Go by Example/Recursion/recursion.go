package main

import "fmt"

func main() {
	fmt.Println("...")
	fmt.Println("fact(7) =", fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fib(n-1) + fib(n-2)
		}
	}
	fmt.Println("fib(7) =", fib(7))
}

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
