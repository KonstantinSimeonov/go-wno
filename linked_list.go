package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
	Length int
}

func (xs *List) Iter () <-chan int {
	ch := make(chan int)
	go func () {
		for x := xs.Head; x != nil; x = x.Next {
			ch <- x.Value
		}

		close(ch)
	} ();
	return ch
}

func NewList() List {
	return List{nil, nil, 0}
}

func l_append_one(l *List, x int) {
	tail := &Node{x,nil}
	l.Length++

	if l.Head == nil {
		l.Head = tail
	} else {
		l.Tail.Next = tail
	}

	l.Tail = tail
}

func AppendMany(l *List, xs ...int) {
	if len(xs) > 0 {
		l_append_one(l, xs[0])
	}

	for i := 1; i < len(xs); i++ {
		l.Tail.Next = &Node{xs[i],nil}
		l.Tail = l.Tail.Next
		l.Length++
	}
}

func node_at(l *List, i int) *Node {
	current := l.Head
	for i > 0 && current != nil {
		current = current.Next
		i--
	}

	return current
}

func At(l *List, i int) *int {
	node := node_at(l, i)
	if node == nil {
		return nil
	}

	return &node.Value
}

func main() {
	xs := NewList()
	AppendMany(&xs, 10, 20, 30, 50)
	fmt.Println(xs)

	for v := range xs.Iter() {
		fmt.Println(v)
	}

	for i := 0; i < xs.Length; i++ {
		fmt.Println(i, *At(&xs, i))
	}
}
