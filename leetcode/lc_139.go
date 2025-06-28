package main

import "fmt"

// 139. 单词拆分

// DictTreeNode 字典树实现，性能并不是很好
type DictTreeNode struct {
	alpha  byte
	isLeaf bool
	elems  map[byte]*DictTreeNode
}
type DictTree struct {
	root *DictTreeNode
}

func NewDictTree() *DictTree {
	return &DictTree{root: &DictTreeNode{alpha: '#', isLeaf: true}}
}

func (t *DictTree) Insert(s string) {
	elem := t.root
	length := len(s)
	for i, v := range s {
		if _, ex := elem.elems[byte(v)]; !ex {
			if elem.elems == nil {
				elem.elems = make(map[byte]*DictTreeNode)
			}
			elem.elems[byte(v)] = &DictTreeNode{
				alpha: byte(v),
			}
		}
		if length-1 == i {
			elem.elems[byte(v)].isLeaf = true
		}
		elem = elem.elems[byte(v)]
	}
}

func (t *DictTree) Find(s string, node *DictTreeNode) bool {

	if len(s) == 0 {
		return true
	}
	for i, v := range s {
		if _, ex := node.elems[byte(v)]; !ex {
			return false
		}
		node = node.elems[byte(v)]
		if node.isLeaf {
			if i == len(s)-1 {
				return true
			} else {
				continue
			}
		}
	}
	return false
}

// 字典树+动态规划， dp[i]表示s的前i个字符是否可以由字典集合组成。判断是否在集合中用到了字典树，其实
// 并不是很有必要，直接使用map判断更简单。dp[i] = true {dp[j] == true && s[j:i]存在字典集合中 && 0 <= j < i}
func wordBreak1(s string, wordDict []string) bool {
	dictTree := NewDictTree()
	res := false
	for _, word := range wordDict {
		dictTree.Insert(word)
	}
	n := len(s)
	dp := make([]bool, n+1)

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i; j >= 0; j-- {
			if dp[j] && dictTree.Find(s[j:i], dictTree.root) {
				dp[i] = true
			}
		}
	}
	res = dp[n]
	return res
}

// 动态规划， dp[i]表示s的前i个字符是否可以由字典集合组成。dp[i] = true {dp[j] == true && s[j:i]存在字典集合中 && 0 <= j < i}
func wordBreak(s string, wordDict []string) bool {

	n := len(s)
	dp := make([]bool, n+1)

	existDict := make(map[string]bool)
	for _, v := range wordDict {
		existDict[v] = true
	}

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i; j >= 0; j-- {
			if dp[j] && existDict[s[j:i]] {
				dp[i] = true
			}
		}
	}
	res := dp[n]
	return res
}

func main() {
	//s := "leetcode"
	//wordDict := []string{"leet", "code"}
	s := "aaaaaaa"
	wordDict := []string{"aaaa", "aa"}
	res := wordBreak(s, wordDict)
	fmt.Println(res)
}
