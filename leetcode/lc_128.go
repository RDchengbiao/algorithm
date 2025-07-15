package main

import "fmt"

// 128. 最长连续序列
// 将所有元素放入map，遍历集合，如果mp[num-1]存在，说明不是最小的，找到连续的最小元素。
// 然后一直遍历该元素为起点的序列最长是多少。 最后更新出最大结果
func longestConsecutive(nums []int) int {
	res := 0
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = true
	}
	for num := range mp {
		if mp[num-1] {
			continue
		}
		cnt := 1
		for i := num + 1; ; i++ {
			if mp[i] {
				cnt++
			} else {
				break
			}
		}
		res = max(res, cnt)
	}
	return res
}
func main() {

	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
