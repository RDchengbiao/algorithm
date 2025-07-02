package main

import "fmt"

// 647. 回文子串
// 枚举回文串的中心，需要考虑以单个字符为中心和以2个字符为中心的情况。然后枚举长度，对比中心两侧。
func countSubstrings(s string) int {

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
func main() {
	s := "aaa"
	res := countSubstrings(s)
	fmt.Println(res)
}
