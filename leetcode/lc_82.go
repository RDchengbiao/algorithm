package main

import "fmt"

//82. 删除排序链表中的重复元素 II

type ListNode struct {
	Val  int
	Next *ListNode
}

// 假头节点，快慢指针，如果是快慢指针所指next节点值不同，则一起向后走，如果是值相同，则一直走快指针，然后将慢指针的next指向快指针的next，
// 向后移动快指针，剔除值相同的节点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := new(ListNode)
	dummyHead.Next = head
	fast := dummyHead.Next
	slow := dummyHead
	for fast != nil && fast.Next != nil {
		for slow.Next != nil && fast.Next != nil && slow.Next.Val != fast.Next.Val {
			fast = fast.Next
			slow = slow.Next
		}
		if slow.Next != nil && fast.Next != nil && slow.Next.Val == fast.Next.Val {
			for slow.Next != nil && fast.Next != nil && slow.Next.Val == fast.Next.Val {
				fast = fast.Next
			}
			slow.Next = fast.Next
			fast = fast.Next
		}
	}

	return dummyHead.Next
}

func main() {
	/*
		node1 := ListNode{
			Val: 1,
		}
		node2 := ListNode{
			Val: 2,
		}
		node3 := ListNode{
			Val: 3,
		}
		node4 := ListNode{
			Val: 3,
		}
		node5 := ListNode{
			Val: 4,
		}
		node6 := ListNode{
			Val: 4,
		}
		node7 := ListNode{
			Val: 5,
		}
		node1.Next = &node2
		node2.Next = &node3
		node3.Next = &node4
		node4.Next = &node5
		node5.Next = &node6
		node6.Next = &node7

	*/

	node1 := ListNode{
		Val: 1,
	}
	node2 := ListNode{
		Val: 1,
	}
	node1.Next = &node2
	head := deleteDuplicates(&node1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
