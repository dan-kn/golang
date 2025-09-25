package main

import (
	"fmt"
	"iter"
	"slices"
)

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

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
