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
	if n, ok := this.m[key]; ok {
		this.Remove(n)
		this.Add(n)
		return n.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.m[key]; ok {
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
		// 若 map > capacity 刪除最後節點
		delete(this.m, this.tail.key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *node) {
	// 將 node prev 設為 null 是由於新加入的 node 要移動到首位
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		// cache 不為空須將 node 放置首位並連結當下 head
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		// cache 為空 tail 為 node
		this.tail = node
	}
}

func (this *LRUCache) Remove(node *node) {
	if this.head == node {
		// 若為首位： node 下個 node 替換成首位
		this.head = node.next
	} else {
		// 非首位：刪除中間 node 串接前個及下個 node
		node.prev.next = node.next
	}
	if this.tail == node {
		// 若為最後：上個 node 即為最後
		this.tail = node.prev
	} else {
		// 若非最後：刪除中間 node 串接前個及下個 node
		node.next.prev = node.prev
	}
}
