package main

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[T]) GetValue() T {
	return n.Value
}
