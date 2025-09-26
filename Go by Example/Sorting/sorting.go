package main

import (
	"fmt"
	"slices"
)

func main() {
	strSlice := []string{"c", "a", "b"}
	slices.Sort(strSlice)

	fmt.Println("strSlice:", strSlice)

	intSlice := []int{3, 1, 2, 6, 4, 5}
	slices.Sort(intSlice)
	fmt.Println("intSlice:", intSlice)

	fmt.Println("strSlice sorted:", slices.IsSorted(strSlice))
	fmt.Println("intSlice sorted:", slices.IsSorted(intSlice))
}
