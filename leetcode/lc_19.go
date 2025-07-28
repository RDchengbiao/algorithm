package main

import "fmt"

//19. 删除链表的倒数第 N 个结点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用一个假头节点方便处理边界情况，然后使用快慢指针，先让快指针走n步，然后快慢指针一起走，直到快指针走到链表结尾，
// 删除慢指针所指向的下一个节点，即可返回假头节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	fast := dummyHead
	slow := dummyHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func main() {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 3,
	}
	node4 := &ListNode{
		Val: 4,
	}
	node5 := &ListNode{
		Val: 5,
	}
	node4.Next = node5
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	//head := removeNthFromEnd(node1, 2)
	//head := removeNthFromEnd(node1, 1)
	head := removeNthFromEnd(node1, 5)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
