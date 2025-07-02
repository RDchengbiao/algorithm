package main

import "fmt"

// 647. 回文子串
// 枚举回文串的中心，需要考虑以单个字符为中心和以2个字符为中心的情况。然后枚举长度，对比中心两侧。
func countSubstrings1(s string) int {

	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		//单字符为中心
		for j := 0; i-j >= 0 && i+j < n; j++ {
			if s[i-j] == s[i+j] {
				res++
			} else {
				break
			}
		}
		//2个字符为中心
		for j := 0; i-j >= 0 && i+1+j < n; j++ {
			if s[i-j] == s[i+j+1] {
				res++
			} else {
				break
			}
		}
	}
	return res
}

// dp 动态规划，dp[i][j]表示s[i]-s[j]是否是回文串。那么只需要s[i]==s[j]且dp[i+1][j-1]是回传串即可。不过需要注意
// j-i小于等于1时，不依赖状态转移方程，同样也是回文串。
func countSubstrings(s string) int {

	n := len(s)
	res := 0
	dp := make(map[int]map[int]bool)
	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if dp[i] == nil {
				dp[i] = make(map[int]bool)
			}
			if dp[i+1] == nil {
				dp[i+1] = make(map[int]bool)
			}
			//i和j所指字母相同，且i和j间距为1或者s[i+1]->s[j-1]子串是回文串
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				res++
				dp[i][j] = true
			}
		}
	}
	return res
}
func main() {
	s := "aaa"
	res := countSubstrings(s)
	fmt.Println(res)
}
