package main

import (
	"fmt"
	"math"
)

//530. 二叉搜索树的最小绝对差

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 由于对搜索二叉树中序遍历出的结果是有序的，那么我们只需要记录前一个节点的值pre，在处理当前节点时计算差值即可，并更新前一个节点值pre为当前值。

func getMinimumDifference(root *TreeNode) int {

	var res int
	res = math.MaxInt
	pre := math.MaxInt / 2
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = min(res, abs(node.Val-pre))
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)

	return res
}
func main() {

	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 0,
	}
	root.Right = &TreeNode{
		Val: 48,
	}
	root.Right.Left = &TreeNode{
		Val: 12,
	}
	root.Right.Right = &TreeNode{
		Val: 49,
	}
	res := getMinimumDifference(root)
	fmt.Println(res)

}
