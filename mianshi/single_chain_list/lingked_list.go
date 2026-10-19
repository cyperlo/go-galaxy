package main

import (
	"fmt"
)

type LinkedList[T any] struct {
	head *Node[T]
}

// 尾插法
func (l *LinkedList[T]) TailAppend(value T) {
	newNode := Node[T]{
		Value: value,
	}
	if l.head == nil {
		l.head = &newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newNode
	}
}

func (l *LinkedList[T]) HeadAppend(value T) {
	newNode := Node[T]{
		Value: value,
	}
	if l.head == nil {
		l.head = &newNode
	} else {
		current := l.head
		newNode.Next = current
		l.head = &newNode
	}
}

func (l *LinkedList[T]) Print() {
	current := l.head

	for current != nil {
		fmt.Printf("%s->", current.GetValue())
		current = current.Next
	}

	fmt.Println("nil")
}
