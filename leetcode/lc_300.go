package main

import "fmt"

// 300. 最长递增子序列

// 动态规划，dp[i]表示以nums[i]结尾的最长递增子序列长度，nums[i]必须取。需要遍历j: 0->i-1之间所有元素，取nums[j]<nums[i]的最大dp[j]+1。
// 时间O(n^2)
func lengthOfLIS(nums []int) int {
	res := 1
	//dp[i] 表示以nums[i]为结尾的最长递增子序列长度
	dp := make(map[int]int)
	dp[0] = 1
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			} else {
				dp[i] = max(dp[i], 1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}
func main() {

	nums := []int{0, 1, 0, 3, 2, 3}
	///nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}
