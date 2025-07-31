package main

// 236. 二叉树的最近公共祖先
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, cur *TreeNode, elems *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	if root == cur {
		*elems = append(*elems, root)
		return true
	}
	leftFind := dfs(root.Left, cur, elems)
	rightFind := dfs(root.Right, cur, elems)
	if leftFind || rightFind {
		*elems = append(*elems, root)
		return true
	}
	return false
}

// 自己想的一个方法是，分别找出从p、q的路径，然后从数组中找到最近公共祖先。
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	pathP := make([]*TreeNode, 0)
	pathQ := make([]*TreeNode, 0)
	dfs(root, p, &pathP)
	dfs(root, q, &pathQ)

	n := len(pathP)
	m := len(pathQ)
	res := root
	for i := 0; i < min(n, m) && pathP[n-1-i] == pathQ[m-1-i]; i++ {
		res = pathP[n-1-i]
	}

	return res
}

// 其实就是2种情况，一种是两个节点分别在某个节点的左右两个子树中。另外一种是，当前节点是p、q中的其中一个，另外一个在该节点的子树中。
// 比较核心的一个逻辑是当左右子树都不是空的时候，说明当前节点就是最近公共祖先。左子树或右子树不是空，那么将结果回溯上去，其中一个最浅
// 的一个就是结果。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	if root == p || root == q {

		return root
	}
	leftSon := lowestCommonAncestor(root.Left, p, q)
	rightSon := lowestCommonAncestor(root.Right, p, q)
	if leftSon != nil && rightSon != nil {
		return root
	}
	if rightSon != nil {
		return rightSon
	}
	return leftSon
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}

	lowestCommonAncestor(root, root.Left, root.Left.Right.Right)
	//lowestCommonAncestor(root, root.Left.Left, root.Left.Right.Right)

}
