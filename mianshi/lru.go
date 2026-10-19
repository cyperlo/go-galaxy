package main

type Node struct {
	key, value int
	pre, next  *Node
}

type LRUCache struct {
	cap        int
	cache      map[int]*Node
	head, tail *Node
}

func NewLRUCache(cap int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head

	return LRUCache{
		cap:   cap,
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (l *LRUCache) Put(key, value int) {
}
