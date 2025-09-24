package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func newPerson(firstName, lastName string, age int) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func main() {
	fmt.Println(Person{"John", "Doe", 32})
	fmt.Println(Person{firstName: "John", lastName: "Doe", age: 32})
	fmt.Println(Person{firstName: "John", lastName: "Doe"})
	fmt.Println(&Person{firstName: "John", lastName: "Doe", age: 32})
	fmt.Println(newPerson("John", "Doe", 32))

	person := Person{firstName: "Danny", lastName: "Kan", age: 999}
	fmt.Println(person)
	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	fmt.Println(person.age)

	sp := &person
	fmt.Println(sp)
	fmt.Println(sp.firstName)
	fmt.Println(sp.lastName)
	fmt.Println(sp.age)

	sp.age = 1000

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Rex",
		isGood: true,
	}

	fmt.Println(dog)
}
