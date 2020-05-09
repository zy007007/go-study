package main

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
//
//示例 1:
//输入: "abab"
//输出: True
//解释: 可由子字符串 "ab" 重复两次构成。
//示例 2:
//输入: "aba"
//输出: False
//示例 3:
//输入: "abcabcabcabc"
//输出: True
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

// 个人思路
// 子串多次构成

import (
	"fmt"
)

// by comment 王小宁。
func repeatedSubstringPattern(s string) bool {
	var str1 string = s + s
	var str2 string = str1[1 : len(str1)-1]

	if strings.Contains(str2, s) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("vim-go")
}

// 假设母串S是由子串s重复N次而成， 则 S+S则有子串s重复2N次， 现在S=Ns， S+S=2Ns 因此S在(S+S)[1:-1]中必出现一次以上
// 评论中这个思路，就是上面代码的实现，想想为什么就一定成立呢
// s+s             s2        s
// abababab     bababa     abab
// abaaba      baab        aba
