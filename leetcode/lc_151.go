package main

import "fmt"

// 151. 反转字符串中的单词

func reverse(s string) string {
	n := len(s)
	ss := []byte(s)
	for i := 0; i <= n-i && i < n; i++ {
		ss[i], ss[n-i-1] = s[n-i-1], s[i]
	}
	return string(ss)
}
func reverseWords(s string) string {
	res := string("")
	news := reverse(s)

	//提取第一个字符串
	tmpStr := ""
	i := 0
	for i = 0; i < len(news); i++ {
		if news[i] == ' ' {
			continue
		} else {
			break
		}
	}

	for ; i < len(news); i++ {
		if news[i] == ' ' {
			break
		}
		tmpStr += string(news[i])
	}
	if tmpStr != "" {
		tmpStr = reverse(tmpStr)
		res = res + tmpStr
	} else {
		return res
	}

	//在提取的字符串前面加空格
	for i < len(news) {
		for ; i < len(news) && news[i] == ' '; i++ {
			continue
		}
		tmpStr = ""
		for ; i < len(news); i++ {
			if news[i] == ' ' {
				break
			}
			tmpStr += string(news[i])
		}
		if tmpStr != "" {
			tmpStr = reverse(tmpStr)
			res = res + " " + tmpStr
		} else {
			break
		}
	}

	return res
}

func main() {
	//s := "  hello world  "
	s := "EPY2giL"
	fmt.Println(reverseWords(s))
}
