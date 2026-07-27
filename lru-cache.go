type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	cap   int
	cache map[int]*Node
	head  *Node // dummy head
	tail  *Node // dummy tail
}

func (this *LRUCache) remove(node *Node) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) insert(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head

	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) moveToFront(node *Node) {
	this.remove(node)
	this.insert(node)
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToFront(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.Val = value
		this.moveToFront(node)
		return
	}

	node := &Node{
		Key: key,
		Val: value,
	}

	this.cache[key] = node
	this.insert(node)

	if len(this.cache) > this.cap {
		lru := this.tail.Prev

		this.remove(lru)
		delete(this.cache, lru.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//  TO1 SOn (2n + 2)