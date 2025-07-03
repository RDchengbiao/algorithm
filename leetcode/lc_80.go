package main

import "fmt"

// 80. 删除有序数组中的重复项 II

// 双指针
func removeDuplicates(nums []int) int {
	n := len(nums)
	i, j, k := 0, 0, 0
	cnt := 0
	for {
		for j < n {
			if nums[i] == nums[j] {
				cnt++
				j++
			} else {
				break
			}
		}
		if cnt >= 2 {
			nums[k] = nums[i]
			k++
			nums[k] = nums[i]
			k++
		} else if cnt == 1 {
			nums[k] = nums[i]
			k++
		}
		cnt = 0
		i = j
		if j == n {
			break
		}
	}
	return k
}

func main() {

	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
