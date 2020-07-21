package main
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	count := 1
	for i:=0; i<len(s)-2;i++{
		if s[i] != s[i+1] {
			count ++
		} else {
			count = 1
		}
	}
	return count
	//"abcabcbb"
	//输出
	//7
	//预期结果
	//3
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int
	n := len(s)
	rk, ans := -1, 0
	for i:= 0; i<n; i++{
		if i != 0 {
			delete(m, s[i])
		}

		for rk + 1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]] ++
			rk ++
		}
		ans = max(ans, rk - i + 1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println("vim-go")
}
