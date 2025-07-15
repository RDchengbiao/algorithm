package main

import "fmt"

// 219. 存在重复元素 II

// 滑动窗口，用map记录窗口内目前的元素都有哪些，同时保证窗口大小始终小于等于k，
// 如果窗口内没有重复元素，则加入，如果有则返回结果。
func containsNearbyDuplicate(nums []int, k int) bool {
	res := false
	n := len(nums)
	filter := map[int]bool{}
	filter[nums[0]] = true
	for i, j := 0, 1; i < j && j < n; {
		if j-i > k {
			delete(filter, nums[i])
			i++
			continue
		}
		if _, ex := filter[nums[j]]; !ex {
			filter[nums[j]] = true
			j++
		} else {
			res = true
			break
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2

	/*
		nums := []int{1, 2, 3, 1}
		k := 3
	*/
	res := containsNearbyDuplicate(nums, k)
	fmt.Println(res)
}
