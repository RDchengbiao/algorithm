package main

import "fmt"

// 54. 螺旋矩阵

// 模拟题，控制好四边边界
func spiralOrder(matrix [][]int) []int {

	res := make([]int, 0)
	xs := 0
	xe := len(matrix) - 1
	ys := 0
	ye := len(matrix[0]) - 1
	x, y := 0, 0

	for len(res) < len(matrix)*len(matrix[0]) {
		for y <= ye && len(res) < len(matrix)*len(matrix[0]) {
			res = append(res, matrix[x][y])
			y++
		}
		y--
		ye--
		for x < xe && len(res) < len(matrix)*len(matrix[0]) {
			x++
			res = append(res, matrix[x][y])
		}
		xe--
		for y > ys && len(res) < len(matrix)*len(matrix[0]) {
			y--
			res = append(res, matrix[x][y])
		}
		ys++
		xs++
		for x > xs && len(res) < len(matrix)*len(matrix[0]) {
			x--
			res = append(res, matrix[x][y])
		}
		x = xs
		y = ys
	}

	return res
}

func main() {
	/*
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
	*/
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	res := spiralOrder(matrix)
	fmt.Println(res)
}
