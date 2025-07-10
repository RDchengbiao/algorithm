package main

import "fmt"

// 209. 长度最小的子数组

// 滑动窗口，左右窗口下标i、j，一直向右移动j直到i->j中的和大于等于target，然后一直向右移动i直到和小于target，记录最小长度。
func minSubArrayLen1(target int, nums []int) int {
	n := len(nums)
	res := 0
	i, j := 0, 0
	sum := 0
	for i <= j && j < n {
		sum += nums[j]
		if sum >= target {
			for i <= j && sum >= target {
				if res != 0 {
					res = min(res, j-i+1)
				} else {
					res = j - i + 1
				}
				sum -= nums[i]
				i++
			}
			j++
		} else {
			j++
		}
	}
	return res
}

// 二分，查找大于等于目标值的最小下标
func binSearch(nums []int, l, r, gap, target int) int {
	res := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target+gap {
			res = mid
			r = mid - 1
		} else if nums[mid] < target+gap {
			l = mid + 1
		}
	}
	return res
}

// 前缀和+二分，只是一种解题思路，不如滑动窗口
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := 0
	prefixSum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			prefixSum[i] = nums[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + nums[i]
		}
	}

	for i := 0; i < n; i++ {
		gap := 0
		if i != 0 {
			gap = prefixSum[i-1]
		}
		idx := binSearch(prefixSum, i, n-1, gap, target)
		if idx == -1 {
			continue
		} else {
			if res == 0 {
				res = idx - i + 1
			} else {
				res = min(res, idx-i+1)
			}
		}
	}
	return res
}

func main() {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	//target := 7
	//nums := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
