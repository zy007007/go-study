package main

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//案例:
//
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.
//
//
//注意事项：您可以假定该字符串只包含小写字母。

// 个人思路
// map 　{l:"0, 1", e:"1, 2", t:"2, 1"}
// 就是操作判断频繁点

import (
	"fmt"
)

func firstUniqChar(s string) int {
	var res [26]int

	for i, ch := range s {
		res[ch-'a'] = i
	}

	for i, ch := range s {
		if i == res[ch-'a'] {
			return i
		} else {
			res[ch-'a'] = -1
		}
	}
	return -1
}

func main() {
	fmt.Println("vim-go")
}
