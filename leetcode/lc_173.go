package main

// 173. 二叉搜索树迭代器

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

type BSTIterator struct {
	stk  *Stack
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		stk: new(Stack),
	}

	it.root = root
	return it
}

func (this *BSTIterator) Next() int {
	for this.root != nil {
		this.stk.Push(this.root)
		this.root = this.root.Left
	}
	root := this.stk.Top()
	this.stk.Pop()
	res := root.Val
	if root.Right != nil {
		this.root = root.Right
	}
	return res
}

func (this *BSTIterator) HasNext() bool {

	return !this.stk.Empty() || this.root != nil
}

func main() {

}
