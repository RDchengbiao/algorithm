package main

import "fmt"

// 101. 对称二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr [][]*TreeNode

func dfs(root *TreeNode, depth int) {

	if root == nil {
		return
	}
	if len(arr) <= depth {
		arr = append(arr, make([]*TreeNode, 0))
	}
	arr[depth] = append(arr[depth], root.Left)
	arr[depth] = append(arr[depth], root.Right)
	dfs(root.Left, depth+1)
	dfs(root.Right, depth+1)
}
func isSymmetric(root *TreeNode) bool {
	arr = make([][]*TreeNode, 0)
	dfs(root, 0)
	for _, subLevel := range arr {
		n := len(subLevel)
		for i := 0; i < len(subLevel); i++ {
			if !(subLevel[i] == nil && subLevel[n-i-1] == nil || (subLevel[i] != nil && subLevel[n-i-1] != nil && subLevel[i].Val == subLevel[n-i-1].Val)) {
				return false
			}
		}
	}
	return true
}

func compare(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return compare(p.Left, q.Right) && compare(p.Right, q.Left)
}

// isSymmetric2 不用多余数组，直接递归比较。递归参数p、q节点，递归比较p左节点和q右节点、比较p右节点和q左节点。
func isSymmetric2(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}

type Queue struct {
	elems []*TreeNode
}

func (q *Queue) Push(elem *TreeNode) {
	q.elems = append(q.elems, elem)
}
func (q *Queue) Pop() *TreeNode {
	if len(q.elems) == 0 {
		return nil
	}
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem
}

func (q *Queue) Empty() bool {
	return len(q.elems) == 0
}

// isSymmetric3 使用迭代实现，首先将根入队2次，每次从队列中取出2个节点p、q，比较值。相同的话，再分别将p左节点和q右节点、p右节点和q左节点分别入队。
func isSymmetric3(root *TreeNode) bool {
	queue := &Queue{}
	queue.Push(root)
	queue.Push(root)
	for !queue.Empty() {
		p := queue.Pop()
		q := queue.Pop()
		if q == nil && p == nil {
			continue
		}
		if q == nil || p == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue.Push(p.Left)
		queue.Push(q.Right)
		queue.Push(p.Right)
		queue.Push(q.Left)
	}
	return true
}

func main() {
	/*
		root := &TreeNode{
			Val: 1,
		}
	*/
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Left = &TreeNode{
		Val: 4,
	}
	root.Right.Right = &TreeNode{
		Val: 3,
	}
	fmt.Println(isSymmetric3(root))

}
