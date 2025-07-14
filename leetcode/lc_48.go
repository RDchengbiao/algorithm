package main

import "fmt"

//官方题解是一次旋转4个位置，将一个点涉及到的旋转位置一次处理完。然后枚举都有哪些点作为起点需要旋转即可。

// 48. 旋转图像
func rotate(matrix [][]int) {

	n := len(matrix)
	//枚举每轮旋转的起始位置
	for x := 0; x < n/2; x++ {
		xx := x
		yy := x
		//枚举哪些列需要旋转
		for j := 0; j <= n-2*(1+x); j++ {
			yy = x + j
			old := matrix[xx][yy]
			newx := xx
			newy := yy
			//经过旋转，最终会回到原点。即：(x,y)->(y,n-x-1) -> (n-x-1,n-y-1) -> (n-y-1,x)->(x,y)
			for {
				newx, newy = newy, n-newx-1
				newv := matrix[newx][newy]
				matrix[newx][newy] = old
				old = newv
				if xx == newx && yy == newy {
					break
				}
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
