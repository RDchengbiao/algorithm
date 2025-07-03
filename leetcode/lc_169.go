package main

import "fmt"

//169. 多数元素

// 消消乐思想，相同加一，不同减一，计数为0时换元素
func majorityElement(nums []int) int {
	n := len(nums)
	elem := nums[0]
	cnt := 1
	for i := 1; i < n; i++ {
		if cnt == 0 {
			elem = nums[i]
			cnt++
			continue
		}
		if nums[i] == elem {
			cnt++
		} else {
			cnt--
		}
	}
	return elem
}
func main() {

	nums := []int{3, 2, 3}
	res := majorityElement(nums)
	fmt.Println(res)
}
