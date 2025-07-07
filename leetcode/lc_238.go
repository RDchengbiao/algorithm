package main

import "fmt"

// 238. 除自身以外数组的乘积

// 维护前缀积，然后从后往前计算结果数组，同时用一个变量维护后缀积。
// 看到题目要求常数空间，没有想到借助结果数组来保存前缀积。蠢爆了。。。
func productExceptSelf(nums []int) []int {

	n := len(nums)
	prefix := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1] * nums[i]
		}
	}

	suffix := nums[n-1]
	for i := 0; i < n; i++ {
		if i == 0 {
			suffix = nums[n-i-1]
			prefix[n-i-1] = prefix[n-i-2]
		} else if i == n-1 {
			prefix[n-i-1] = suffix
		} else {
			prefix[n-i-1] = prefix[n-i-2] * suffix
			suffix = suffix * nums[n-i-1]
		}
	}
	return prefix
}

func main() {

	nums := []int{-1, 1, 0, -3, 3}
	//nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
