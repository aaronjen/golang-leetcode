package main

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start

// LRUNode struct
type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

// LRUCache struct
type LRUCache struct {
	table map[int]*LRUNode
	head  *LRUNode
	tail  *LRUNode
	cap   int
}

// Constructor for LRUCache
func Constructor(capacity int) LRUCache {
	head := new(LRUNode)
	tail := new(LRUNode)
	head.next = tail
	tail.prev = head

	return LRUCache{
		head:  head,
		tail:  tail,
		cap:   capacity,
		table: map[int]*LRUNode{},
	}
}

func (c *LRUCache) removeNode(n *LRUNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) insertNode(n *LRUNode) {
	c.head.next.prev = n
	n.next = c.head.next
	c.head.next = n
	n.prev = c.head
}

// Get for LRUCache
func (c *LRUCache) Get(key int) int {
	if node, ok := c.table[key]; ok {
		c.removeNode(node)
		c.insertNode(node)
		return node.value
	}

	return -1
}

// Put for LRUCache
func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.table[key]; ok {
		node.value = value
		c.removeNode(node)
		c.insertNode(node)
	} else {
		node := new(LRUNode)
		node.value = value
		node.key = key
		if len(c.table) == c.cap {
			last := c.tail.prev
			c.removeNode(last)
			delete(c.table, last.key)
		}
		c.insertNode(node)
		c.table[key] = node
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
