package main

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
//你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
//
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
//示例 1:
//s = "abc", t = "ahbgdc"
//
//返回 true.
//
//示例 2:
//s = "axc", t = "ahbgdc"
//
//返回 false.
//
//后续挑战 :
//
//如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

// 个人思路
// map，还是map呗
// 结果看，还有其他

import (
	"fmt"
)

// by comment a316523235
func isSubsequence(s string, t string) bool {
	return methodPoint(s, t)
}

func methodPoint(s, t string) bool {
	if len(s) > len(t) {
		return false
	}

	t1, t2 := 0, 0
	for t1 < len(s) && t2 < len(t2) {
		if s[t1] == t[t2] {
			t1++
		}
		t2++
	}

	return t1 == len(s)
}

func main() {
	fmt.Println("vim-go")
}
