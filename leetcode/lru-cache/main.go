package main

type LRUCache struct {
	head *node
	tail *node

	cap  int
	dict map[int]*node
}

func Constructor(capacity int) LRUCache {
	head, tail := newNode(-1, -1), newNode(-1, -1)
	head.next = tail
	tail.prev = head
	return LRUCache{
		head: head,
		tail: tail,
		cap:  capacity,
		dict: make(map[int]*node),
	}
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.dict[key]
	if !ok {
		return -1
	}

	c.remove(node)
	c.add(node)

	return node.val
}

func (c *LRUCache) add(n *node) {
	prevEnd := c.tail.prev
	prevEnd.next = n
	n.prev = prevEnd
	n.next = c.tail
	c.tail.prev = n
}

func (c *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) Put(key int, value int) {
	if oldNode, ok := c.dict[key]; ok {
		c.remove(oldNode)
	}

	n := newNode(key, value)
	c.dict[key] = n
	c.add(n)

	if len(c.dict) > c.cap {
		nodeToDelete := c.head.next
		c.remove(nodeToDelete)
		delete(c.dict, nodeToDelete.key)
	}
}

type node struct {
	val  int
	key  int
	prev *node
	next *node
}

func newNode(key, val int) *node {
	return &node{
		key: key,
		val: val,
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
