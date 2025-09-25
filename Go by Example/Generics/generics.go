package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for j := range s {
		if v == s[j] {
			return j
		}
	}
	return -1
}

type List[T any] struct{ head, tail *elem[T] }

type elem[T any] struct {
	val  T
	next *elem[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &elem[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &elem[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	s := []string{"a", "b", "c"}

	fmt.Println("index of 'b':", SlicesIndex(s, "b"))

	_ = SlicesIndex[[]string, string](s, "b")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Println("list:", lst.AllElements())
}
