package main

import "fmt"

func main() {
	intArr := [10]int{}
	for j := range intArr {
		intArr[j] = j
	}

	fmt.Println("intArr:", intArr)

	sum := 0
	for _, val := range intArr {
		sum += val
	}

	fmt.Println("sum:", sum)

	for j, val := range intArr {
		if val == 3 {
			fmt.Println("index:", j, "|", "value:", val)
		}
	}

	kvs := map[string]string{"a": "apple", "d": "danny"}
	for key, val := range kvs {
		fmt.Printf("%s -> %s\n", key, val)
	}

	for key := range kvs {
		fmt.Println("key:", key)
	}

	for j, val := range "golang" {
		fmt.Println(j, val)
	}
}
