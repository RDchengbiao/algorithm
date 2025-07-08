package main

import "fmt"

//11. 盛最多水的容器

// 贪心，双指针。移动低的一边，才有可能是结果更大，移动高的一边，对结果不会有任何好处。
func maxArea(height []int) int {
	res := 0
	i, j := 0, len(height)-1
	for i < j {
		res = max(res, (j-i)*min(height[j], height[i]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}
