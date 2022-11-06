package main

type LRUCache struct {
	cache 		map[int]*DLinkedNode
	head		*DLinkedNode
	tail 		*DLinkedNode
	size 		int
	capacity 	int
}

type DLinkedNode struct {
	key		int
	value 	int
	prev	*DLinkedNode
	next 	*DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: 		make(map[int]*DLinkedNode),
		head: 		&DLinkedNode{},
		tail: 		&DLinkedNode{},
		capacity: 	capacity,
		size:		0,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if !ok {
		node := &DLinkedNode{key: key, value: value}
		this.cache[key] = node
		this.addToHead(node)
		this.size += 1
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size -= 1
		}
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	
}