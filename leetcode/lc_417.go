package main

import "fmt"

//417. 太平洋大西洋水流问题

//dfs 逆流而上，记录两种大洋流过的状态，使用一个map记录即可，区分是哪种洋流流经，或者是2种都流经状态

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(heights [][]int, x, y int, cntMp map[int]map[int]int, typ int) {
	if _, ex := cntMp[x]; !ex {
		cntMp[x] = make(map[int]int)
	}
	if cntMp[x][y] == 0 {
		if typ == 0 {
			cntMp[x][y] = 1
		} else {
			cntMp[x][y] = 2
		}
	} else if cntMp[x][y] == 1 {
		if typ == 0 {
			return
		} else if typ == 1 {
			cntMp[x][y] = 3
		}
	} else if cntMp[x][y] == 2 {
		if typ == 0 {
			cntMp[x][y] = 3
		} else {
			return
		}
	} else if cntMp[x][y] == 3 {
		return
	}

	for i := 0; i < len(directions); i++ {
		newx := x + directions[i][0]
		newy := y + directions[i][1]
		if newx < 0 || newy < 0 || newx >= len(heights) || newy >= len(heights[0]) {
			continue
		}
		if heights[newx][newy] < heights[x][y] {
			continue
		}
		if _, ex := cntMp[newx]; !ex {
			cntMp[newx] = make(map[int]int)
		}
		dfs(heights, newx, newy, cntMp, typ)
	}
}
func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)
	n := len(heights)
	m := len(heights[0])
	cntMp := map[int]map[int]int{}
	for i := 0; i < n; i++ {
		dfs(heights, i, 0, cntMp, 0)
	}
	for i := 0; i < m; i++ {
		dfs(heights, 0, i, cntMp, 0)
	}
	for i := 0; i < n; i++ {
		dfs(heights, i, m-1, cntMp, 1)
	}
	for i := 0; i < m; i++ {
		dfs(heights, n-1, i, cntMp, 1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if cntMp[i][j] == 3 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func main() {

	heights := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	res := pacificAtlantic(heights)
	fmt.Println(res)
}
