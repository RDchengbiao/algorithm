package main

import "fmt"

// 117. 填充每个节点的下一个右侧节点指针 II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 最容易想到的方法是层序遍历，可以使用队列将每层的节点存入其中，每层的数量可以通过队列长度n获取到，然后取出n个节点，同时将下层节点加入队列，同时将下层节点串接起来形成链表，
// 不过使用队列空间无法做到常数。仔细观察可以发现，在遍历第x层时，第x+1层已经被连接成链表，可以通过遍历链表取代使用队列保存层序关系。那么需要记录下层的第一个节点作为链表头
// 节点，同时需要一个指针将下层串成一个链表。

func connect(root *Node) *Node {

	cur := root
	nextLevelHead := (*Node)(nil)
	nextLevelPre := (*Node)(nil)
	for cur != nil {
		if cur.Left != nil && nextLevelHead == nil {
			nextLevelHead = cur.Left
		}
		if cur.Right != nil && nextLevelHead == nil {
			nextLevelHead = cur.Right
		}
		if nextLevelPre == nil {
			nextLevelPre = nextLevelHead
		}
		if nextLevelPre != nil {
			if cur.Left != nil && nextLevelPre != cur.Left {
				nextLevelPre.Next = cur.Left
				nextLevelPre = cur.Left
			}
			if cur.Right != nil && nextLevelPre != cur.Right {
				nextLevelPre.Next = cur.Right
				nextLevelPre = cur.Right
			}
		}

		//遍历链表
		cur = cur.Next
		//换层
		if cur == nil {
			cur = nextLevelHead
			nextLevelHead = nil
			nextLevelPre = nil
		}
	}

	return root

}

func main() {

	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Right = &Node{Val: 7}

	/*
		root := &Node{Val: 1}
		root.Left = &Node{Val: 2}

	*/
	connect(root)
	fmt.Println(root.Val)
}
