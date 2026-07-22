package medium

type CacheNode struct {
	val  int
	key  int
	next *CacheNode
	prev *CacheNode
}

type LRUCache struct {
	capacity int
	data     map[int]*CacheNode
	head     *CacheNode
	tail     *CacheNode
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*CacheNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		this.relink(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		node.val = value
		this.relink(node)
		return
	}
	node := &CacheNode{val: value, key: key}
	if len(this.data) == this.capacity {
		this.evictTail()
	}
	this.data[key] = node
	this.relink(node)
}

func (c *LRUCache) relink(node *CacheNode) {
	if node == c.head {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} 
	if node == c.tail && node.prev != nil {
		c.tail = node.prev
	}
	node.prev = nil
	node.next = c.head
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node
	if c.tail == nil {
		c.tail = node
	}
}

func (this *LRUCache) evictTail() {
	delete(this.data, this.tail.key)
	if this.tail.prev != nil {
		this.tail, this.tail.next = this.tail.prev, nil
	} else {
		this.tail, this.head = nil, nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
