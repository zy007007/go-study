package main

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//说明:
//你可以假设字符串只包含小写字母。
//
//进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

// 个人思路
// 长度一样，排序之后遍历对比

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	nums := [26]int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		nums[s[i]-'a']++
		nums[t[i]-'a']--
	}

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
