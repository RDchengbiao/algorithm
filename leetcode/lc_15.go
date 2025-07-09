package main

import (
	"fmt"
	"sort"
)

// 15. 三数之和
func binSearch(nums []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// 枚举两个数，二分查找第三个数
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	filter := make(map[int]map[int]bool)
	res := make([][]int, 0)
	//枚举单个数，然后双指针选取另外两个值
	n := len(nums)
	for i := 0; i < n; i++ {

		j := i + 1
		for ; j < n; j++ {
			aim := -(nums[i] + nums[j])
			ret := binSearch(nums, j+1, n-1, aim)
			if filter[nums[i]] == nil {
				filter[nums[i]] = make(map[int]bool)
			}
			if ret != -1 {
				if filter[nums[i]][nums[j]] {
					continue
				}
				filter[nums[i]][nums[j]] = true
				res = append(res, []int{nums[i], nums[j], aim})
			}
		}
	}
	return res
}

// 排序+双指针。 枚举第一个数，双指针O(n)复杂度计算两数之和，通过过滤重复元素去重
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		//对枚举的i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		aim := -nums[i]
		for j < k {
			if nums[j]+nums[k] > aim {
				k--
			} else if nums[j]+nums[k] < aim {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				//对j位置去重
				if nums[j] == nums[j-1] {
					for j < k && nums[j] == nums[j-1] {
						j++
					}
				}
			}
		}
	}
	return res
}
func main() {

	nums := []int{-1, 0, 1, 0}
	//nums := []int{0, 0, 0}
	//nums := []int{0, 1, 1}
	//nums := []int{-2, 0, 1, 1, 2}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
