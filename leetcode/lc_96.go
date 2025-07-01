package main

import "fmt"

// 动态规划，dp[i]表示当有i个节点时有多少种二叉搜索树。

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		//枚举j表示将j作为根节点，左子树有j-1个节点和右子树有i-j个节点时，一共有多少种组合
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	res := numTrees(4)
	fmt.Println(res)
}
