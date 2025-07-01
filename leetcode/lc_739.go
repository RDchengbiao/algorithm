package main

import "fmt"

// 739. 每日温度

type Temperature struct {
	pos int
	val int
}

// MonotonicityStack 单调栈，从后向前遍历，栈底维护最大的值，栈顶最小。当新元素比栈顶小，计算两个元素坐标差，同时入栈。当新元素大于等于栈顶，则
// 弹出栈顶。最后将结果数组翻转一下即可。
type MonotonicityStack struct {
	elems []Temperature
}

func (d *MonotonicityStack) Pop() Temperature {
	temp := d.elems[len(d.elems)-1]
	d.elems = d.elems[:len(d.elems)-1]
	return temp
}

func (d *MonotonicityStack) Push(temp Temperature) {
	d.elems = append(d.elems, temp)
}
func (d *MonotonicityStack) Len() int {
	return len(d.elems)
}
func (d *MonotonicityStack) Top() Temperature {
	return d.elems[len(d.elems)-1]
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, 0)
	stack := MonotonicityStack{}
	n := len(temperatures)
	for i := n - 1; i >= 0; i-- {
		temp := Temperature{
			pos: i,
			val: temperatures[i],
		}

		for {
			if stack.Len() > 0 && stack.Top().val <= temp.val {
				stack.Pop()
				continue
			}
			if stack.Len() == 0 {
				stack.Push(temp)
				res = append(res, 0)
				break
			}
			if stack.Top().val > temp.val {
				res = append(res, stack.Top().pos-temp.pos)
				stack.Push(temp)
				break
			}
		}
	}
	for i := 0; i <= n-i-1; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}

func main() {

	temperatures := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	res := dailyTemperatures(temperatures)
	fmt.Println(res)
}
