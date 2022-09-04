package main

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	// TODO: implement here
}

type LRUCache struct {
	head     *node
	tail     *node
	m        map[int]*node
	capacity int
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{m: make(map[int]*node), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if ok {
		this.Remove(n)
		this.Add(n)
		return n.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.m[key]
	if ok {
		n.value = value
		this.Remove(n)
		this.Add(n)
		return
	} else {
		n = &node{key: key, value: value}
		this.m[key] = n
		this.Add(n)
	}
	if len(this.m) > this.capacity {
		delete(this.m, this.tail.key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) Remove(node *node) {
	if this.head == node {
		// when node is head
		this.head = node.next
	} else {
		node.prev.next = node.next
	}
	if this.tail == node {
		// when node is tail
		this.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
}
