package main

import "fmt"

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func countAlive(board [][]int, x, y int) int {
	aliveCount := 0
	n := len(board)
	m := len(board[0])
	for i := 0; i < len(directions); i++ {
		nx := x + directions[i][0]
		ny := y + directions[i][1]
		if nx >= 0 && nx < n && ny >= 0 && ny < m {
			if board[nx][ny] == 1 {
				aliveCount++
			} else if board[nx][ny] == 3 {
				aliveCount++
			}
		}
	}
	return aliveCount
}

// 289. 生命游戏
func gameOfLife(board [][]int) {

	n := len(board)
	m := len(board[0])
	//2表示本来是0，需要设置为1. 3表示本来是1，需要设置为0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			aliveCnt := countAlive(board, i, j)
			if board[i][j] == 0 {
				if aliveCnt == 3 {
					board[i][j] = 2
				}
			} else if board[i][j] == 1 {
				if aliveCnt < 2 || aliveCnt > 3 {
					board[i][j] = 3
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	gameOfLife(board)
	fmt.Println(board)
}
