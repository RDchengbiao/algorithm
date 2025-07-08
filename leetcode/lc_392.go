package main

import "fmt"

// 392. 判断子序列

// 动态规划，dp[i][j]表示字符串t从i位置第一次出现j字符的位置
func isSubsequence1(s string, t string) bool {
	n := len(s)
	m := len(t)
	dp := make(map[int]map[byte]int)

	for i := m; i >= 0; i-- {
		if _, ex := dp[i]; !ex {
			dp[i] = make(map[byte]int)
		}
		if i == m {
			for j := 0; j < 26; j++ {
				alpha := byte(j + 'a')
				dp[i][alpha] = m
			}
			continue
		}
		for j := 0; j < 26; j++ {
			alpha := byte(j + 'a')
			dp[i][alpha] = dp[i+1][alpha]
		}
		dp[i][t[i]] = i
	}

	j := 0
	for i := 0; i < n; i++ {
		alpha := s[i]
		if dp[j][alpha] == m {
			return false
		}
		j = dp[j][alpha] + 1
	}
	return true
}

// 贪心，双指针
func isSubsequence(s string, t string) bool {
	n := len(s)
	m := len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == n {
		return true
	}
	return false
}
func main() {

	s := "adc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
