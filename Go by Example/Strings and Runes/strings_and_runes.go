package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const str = "something goes here..."

	fmt.Println("len:", len(str))

	for j := range str {
		fmt.Printf("%x ", str[j])
	}

	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(str))

	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString ->")
	for j, w := 0, 0; j < len(str); j += w {
		runeValue, width := utf8.DecodeRuneInString(str[j:])
		fmt.Printf("%#U starts at %d\n", runeValue, j)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	switch r {
	case 's':
		fmt.Println("found 's'")
	case 'h':
		fmt.Println("found 'h'")
	}
}
