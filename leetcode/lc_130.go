package main

import "fmt"

//130. 被围绕的区域

var directions = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func dfs(board [][]byte, x, y int, vis map[int]map[int]bool) {
	vis[x][y] = true
	for _, direction := range directions {
		nx := x + direction[0]
		ny := y + direction[1]
		if nx < 0 || nx >= len(board) || ny < 0 || ny >= len(board[0]) {
			continue
		}
		if vis[nx] == nil {
			vis[nx] = make(map[int]bool)
		}

		if !vis[nx][ny] && board[nx][ny] == 'O' {
			dfs(board, nx, ny, vis)
		}
	}
}

//很难判断哪些O是被X包围的，但是很容易知道哪些O是不能被包围的，所以我们只需要从四周开始找不能被包围的O即可，最后没有访问过的就是可以被X
//包围的，更新为X即可。

func solve(board [][]byte) {

	vis := make(map[int]map[int]bool)
	n := len(board)
	m := len(board[0])
	for i := 0; i < n; i++ {
		if vis[i] == nil {
			vis[i] = make(map[int]bool)
		}
		if !vis[i][0] && board[i][0] == 'O' {
			dfs(board, i, 0, vis)
		}
		if !vis[i][m-1] && board[i][m-1] == 'O' {
			dfs(board, i, m-1, vis)
		}
	}
	for j := 0; j < m; j++ {
		if !vis[0][j] && board[0][j] == 'O' {
			dfs(board, 0, j, vis)
		}
		if !vis[n-1][j] && board[n-1][j] == 'O' {
			dfs(board, n-1, j, vis)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' && !vis[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}

	/*

		board := [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			}
	*/

	solve(board)
	fmt.Println(board)
}
