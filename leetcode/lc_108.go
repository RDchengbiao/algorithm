package main

import "fmt"

//108. 将有序数组转换为二叉搜索树
//其实题目要求是平衡二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 由于是要求平衡二叉搜索树，那么从中间划分，然后左右两侧分别递归操作即可。

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	idx := len(nums) / 2
	left := sortedArrayToBST(nums[:idx])
	right := sortedArrayToBST(nums[idx+1:])
	return &TreeNode{Val: nums[idx], Left: left, Right: right}
}
func main() {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)
	fmt.Println(res)

}
