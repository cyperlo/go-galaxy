package main

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// 尾插法
func (l *DoublyLinkedList) Append(val int) {
	newNode := &Node{
		Val: val,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	l.Tail.Next = newNode
	newNode.Prev = l.Head
	l.Tail = newNode
}

// 头插法
func (l *DoublyLinkedList) PreAppend(val int) {
	newNode := &Node{
		Val: val,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
}

func (l *DoublyLinkedList) Delete(node *Node) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}
}
