package main

import "fmt"

//94. 二叉树的中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode

func (s *Stack) Push(n *TreeNode) {
	*s = append(*s, n)
}
func (s *Stack) Pop() *TreeNode {
	if len(*s) == 0 {
		return nil
	}
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}
func (s *Stack) Top() *TreeNode {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// 使用迭代模拟递归中序遍历。
// 一直迭代左子树，将节点一直压栈。出栈打印节点的值。如果节点右子树不为空，移动指针到右子树，
// 继续迭代左子树，并重复这个过程。

func inorderTraversal(root *TreeNode) []int {

	stk := &Stack{}
	res := make([]int, 0)
	for !stk.Empty() || root != nil {
		for root != nil {
			stk.Push(root)
			root = root.Left
		}
		rootTmp := stk.Top()
		stk.Pop()
		res = append(res, rootTmp.Val)
		if rootTmp.Right != nil {
			root = rootTmp.Right
		}
	}
	return res
}

func main() {

	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 4}
	res := inorderTraversal(root)
	fmt.Println(res)

}
