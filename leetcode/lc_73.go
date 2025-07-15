package main

//73. 矩阵置零

//题目的难点在于怎么不开辟多余空间，记录哪些行和列需要设置0.可以用数组本身的第一行和
//第一列来记录，哪些列和哪些行需要设置0.同时用两个变量记录第一行和第一列是否本身包含0.最后遍历
//第一行，如果是0，那么将该列设置0；遍历第一列，如果是0，那么将该行设置0.最后根据变量记录的行和列变量
// 将第一行和第一列全部设置0.

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	firstRowHasZero := false
	firstColHasZero := false

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && matrix[i][j] == 0 {
				firstRowHasZero = true
			}
			if j == 0 && matrix[i][j] == 0 {
				firstColHasZero = true
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < m; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < n; j++ {
				matrix[j][i] = 0
			}
		}
	}
	if firstRowHasZero {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColHasZero {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

}

func main() {

	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	/*
		matrix := [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			}

	*/
	setZeroes(matrix)
}
