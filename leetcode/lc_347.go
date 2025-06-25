package main

import "fmt"

// 出现的数字统计计数，按照桶计数的思想，将出现次数相同的数字维护起来，之后提取前k个即可
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	n := len(nums)
	times := make(map[int]int)
	countList := make(map[int][]int)
	for _, num := range nums {
		times[num]++
	}
	for num, count := range times {
		countList[count] = append(countList[count], num)
	}
	for i := n; i >= 1; i-- {
		for j := 0; j < len(countList[i]); j++ {
			if k <= len(res) {
				return res
			} else {
				res = append(res, countList[i][j])
			}
		}
	}
	return res
}
func main() {

	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	res := topKFrequent(nums, k)
	fmt.Println(res)
}
