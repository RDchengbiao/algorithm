package main

import "fmt"

//114. 二叉树展开为链表

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 找到左子树最靠近右子树的节点，最靠近右子树就是优先顺序为右子树、左子树、当前节点，并把该节点的右指针指向右子树的根节点。

func dfs(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	leftTree := dfs(root.Left)
	rightTree := dfs(root.Right)
	if leftTree != nil {
		leftTree.Right = root.Right
	}
	if root.Left != nil {
		root.Right = root.Left
	}
	root.Left = nil

	if rightTree != nil {
		return rightTree
	}
	if leftTree != nil {
		return leftTree
	}
	return root
}
func flatten(root *TreeNode) {
	dfs(root)
}

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}

	flatten(root)
	fmt.Println(root.Val)
}
