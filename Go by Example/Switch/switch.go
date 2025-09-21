package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range 10 {
		fmt.Println("write", i, "as")
		switch i {
		case 1:
			fmt.Println("--- ONE ---")
		case 2:
			fmt.Println("--- TWO ---")
		case 3:
			fmt.Println("--- THREE ---")
		default:
			fmt.Println("OTHER")
		}
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i'm a bool")
		case int:
			fmt.Println("i'm an int")
		default:
			fmt.Printf("don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(false)

	whatAmI(10)

	whatAmI("something goes here...")
}
