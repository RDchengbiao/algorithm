package main

import "fmt"

//138. 随机链表的复制

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func NewNode(val int) *Node {
	node := new(Node)
	node.Val = val
	return node
}
func copyRandomList(head *Node) *Node {

	//restructure list, insert new node
	tmpHead := head
	var copyRandomHead *Node
	for tmpHead != nil {
		copyListNode := NewNode(tmpHead.Val)
		if copyRandomHead == nil {
			copyRandomHead = copyListNode
		}
		copyListNode.Next = tmpHead.Next
		copyListNode.Random = tmpHead.Random
		tmpHead.Next = copyListNode
		tmpHead = tmpHead.Next.Next
	}
	//fix new node's random pointer
	tmpHead = head
	for tmpHead != nil {
		tmpCopyNode := tmpHead.Next
		if tmpCopyNode.Random != nil {
			tmpCopyNode.Random = tmpCopyNode.Random.Next
		}
		tmpHead = tmpCopyNode.Next
	}
	//take list and copy list apart
	//A->a->B->b
	tmpHead = head
	for tmpHead != nil {
		tmpCopyNode := tmpHead.Next
		tmpHead.Next = tmpCopyNode.Next
		if tmpHead.Next != nil {
			tmpCopyNode.Next = tmpHead.Next.Next
		}
		tmpHead = tmpHead.Next
	}
	return copyRandomHead
}
func main() {
	newNode7 := NewNode(7)
	newNode13 := NewNode(13)
	newNode11 := NewNode(11)
	newNode10 := NewNode(10)
	newNode1 := NewNode(1)
	newNode7.Next = newNode13
	newNode13.Next = newNode11
	newNode13.Random = newNode7
	newNode11.Next = newNode10
	newNode11.Random = newNode1

	newNode10.Next = newNode1
	newNode10.Random = newNode11
	newNode1.Random = newNode7

	newRandomNode := copyRandomList(newNode7)
	fmt.Println(newRandomNode)
}
