package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
