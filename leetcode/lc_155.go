package main

import "fmt"

// 155. 最小栈

// Node 在存储的节点中，增加最小值字段，在入栈的时候，判断当前值和栈顶元素哪个更小，维护一个最小值即可。
type Node struct {
	val    int
	minVal int
}
type MinStack struct {
	elems []Node
}

func Constructor() MinStack {
	obj := MinStack{}
	return obj
}

func (this *MinStack) Push(val int) {
	newElem := Node{val: val, minVal: val}
	if len(this.elems) != 0 {
		if this.elems[len(this.elems)-1].minVal < val {
			newElem.minVal = this.elems[len(this.elems)-1].minVal
		}
	}
	this.elems = append(this.elems, newElem)
}

func (this *MinStack) Pop() {
	this.elems = this.elems[:len(this.elems)-1]
}

func (this *MinStack) Top() int {
	top := this.elems[len(this.elems)-1]
	return top.val
}

func (this *MinStack) GetMin() int {
	top := this.elems[len(this.elems)-1]
	return top.minVal
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param_4 := obj.GetMin()
	fmt.Println(param_4)
	obj.Pop()
	param_3 := obj.Top()
	fmt.Println(param_3)
	param_4 = obj.GetMin()
	fmt.Println(param_4)
}
