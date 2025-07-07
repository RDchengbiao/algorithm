package main

import "fmt"

// 189. 轮转数组

// 从数组起始下标0开始，向右移动k个位置，提出原位置的元素，继续向右迭代处理。如果右移k个位置回到了0下标，则递增下标，继续迭代。
// 直到处理n次结束
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	pos := 0
	cnt := 0
	for pos < n && cnt < n {
		start := pos
		oldVal := nums[pos]
		for cnt < n {
			newVal := nums[(pos+k)%n]
			nums[(pos+k)%n] = oldVal
			pos = (pos + k) % n
			oldVal = newVal
			cnt++
			if pos == start {
				pos++
				break
			}
		}
	}
}
func main() {

	/*
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
	*/
	nums := []int{1, 2, 3, 4}
	k := 2
	rotate(nums, k)
	fmt.Println(nums)
}
