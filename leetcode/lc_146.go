package main

import "fmt"

//146. LRU 缓存

// LRUCache LRU缓存
type LRUCache struct {
	capacity int
	head     *DoubleLinkNode
	tail     *DoubleLinkNode
	hashmap  map[int]*DoubleLinkNode
}
type DoubleLinkNode struct {
	next *DoubleLinkNode
	prev *DoubleLinkNode
	val  int
	key  int
}

func newDoubleLinkNode(key, val int) *DoubleLinkNode {
	return &DoubleLinkNode{
		key: key,
		val: val,
	}
}

func (this *LRUCache) moveHead(node *DoubleLinkNode) {
	//移动链表中的节点到头部
	pre := node.prev
	next := node.next
	pre.next = next
	next.prev = pre
	oldHead := this.head.next
	this.head.next = node
	node.next = oldHead
	node.prev = this.head
	oldHead.prev = node
}
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		hashmap:  make(map[int]*DoubleLinkNode),
		head:     newDoubleLinkNode(0, 0),
		tail:     newDoubleLinkNode(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if this.hashmap == nil {
		return -1
	}
	if _, ex := this.hashmap[key]; !ex {
		return -1
	} else {
		//move current node to head
		this.moveHead(this.hashmap[key])
		return this.hashmap[key].val
	}

}

func (this *LRUCache) Put(key int, value int) {

	if _, ex := this.hashmap[key]; ex {
		this.hashmap[key].val = value
		this.moveHead(this.hashmap[key])
	} else {
		if len(this.hashmap) >= this.capacity {
			cur := this.tail.prev
			cur.prev.next = this.tail
			this.tail.prev = cur.prev
			delete(this.hashmap, cur.key)
		}

		//移动新节点到链表头部
		newNode := newDoubleLinkNode(key, value)
		this.hashmap[key] = newNode
		oldHead := this.head.next
		this.head.next = this.hashmap[key]
		this.hashmap[key].next = oldHead
		this.hashmap[key].prev = this.head
		oldHead.prev = this.hashmap[key]
	}
}

func main() {

	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)           // 缓存是 {1=1}
	lRUCache.Put(2, 2)           // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3)) // 返回 3
	fmt.Println(lRUCache.Get(4)) // 返回 4

}
