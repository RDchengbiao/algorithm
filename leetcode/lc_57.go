package main

import "fmt"

// 57. 插入区间

// 分析3种情况，1. 区间end小于和插入区间的start，两者不重叠 2. 区间start大于插入区间的end，两者不重叠 3.区间和插入区间重叠
// 对情况1不重叠的情况，直接放入结果数组。对于重叠的情况，需要记录新区间的左边界和右边界，左边界是最小值，右边界是最大值。对情况2
// 不重叠的情况，第一次时，需要将情况3保存的结果放入结果数组。需要特殊处理区间数组为空的情况。
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	newStart := newInterval[0]
	newEnd := newInterval[1]
	flag := false
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else if interval[0] > newInterval[1] {
			if !flag {
				flag = true
				res = append(res, []int{newStart, newEnd})
			}
			res = append(res, interval)
		} else {
			newStart = min(newStart, interval[0])
			newEnd = max(newEnd, interval[1])
		}
	}
	if !flag {
		flag = true
		res = append(res, []int{newStart, newEnd})
	}
	return res
}

func main() {

	intervals := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	newInterval := []int{4, 8}
	res := insert(intervals, newInterval)
	fmt.Println(res)
}
