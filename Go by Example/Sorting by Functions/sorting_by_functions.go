package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	lenCmp := func(a, b string) int { return cmp.Compare(len(a), len(b)) }

	slices.SortFunc(fruits, lenCmp)
	fmt.Println("fruits:", fruits)

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	people := []Person{
		{firstName: "John", lastName: "Doe", age: 32},
		{firstName: "Jane", lastName: "Doe", age: 18},
		{firstName: "Danny", lastName: "Smith", age: 100},
	}

	ageCmp := func(a, b Person) int { return cmp.Compare(a.age, b.age) }

	slices.SortFunc(people, ageCmp)
	fmt.Println("people:", people)
}
