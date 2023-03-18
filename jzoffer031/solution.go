// https://leetcode-cn.com/problems/lru-cache/
package jzoffer031

type LRUCache struct {
	head     *Node
	tail     *Node
	m        map[int]*Node
	capacity int
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.m = make(map[int]*Node)

	cache.head = &Node{
		Key: -1,
		Val: -1,
	}
	cache.tail = &Node{
		Key: -1,
		Val: -1,
	}
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head

	cache.capacity = capacity
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}

	this.moveToTail(node, node.Val)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.moveToTail(node, value)
	} else {
		if len(this.m) == this.capacity {
			toDeleted := this.head.Next
			this.deleteNode(toDeleted)
			delete(this.m, toDeleted.Key)
		}

		node := &Node{
			Key: key,
			Val: value,
		}
		this.insertToTail(node)
		this.m[key] = node
	}

}

func (this *LRUCache) moveToTail(node *Node, newVal int) {
	this.deleteNode(node)

	node.Val = newVal
	this.insertToTail(node)
}

func (this *LRUCache) deleteNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) insertToTail(node *Node) {
	this.tail.Prev.Next = node
	node.Prev = this.tail.Prev
	node.Next = this.tail
	this.tail.Prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
