package main

//给定两个字符串 s 和 t，判断它们是否是同构的。
//
//如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//
//所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
//
//示例 1:
//
//输入: s = "egg", t = "add"
//输出: true
//示例 2:
//
//输入: s = "foo", t = "bar"
//输出: false
//示例 3:
//
//输入: s = "paper", t = "title"
//输出: true
//说明:
//你可以假设 s 和 t 具有相同的长度。

// 个人思路
// 同构异分体
// 果然是map

import "fmt"

// by mojocn
func isIsomorphiComment(s string, t string) bool {
	v1 := oneWay(s, t)
	v2 := oneWay(t, s)
	return v1 && v2
}

func oneWay(s, t string) bool {
	tmp := map[int32]int32{}

	for idx, sr := range s {
		k := int32[t[idx]]
		v := int32(sr)

		data, ok := tmp[k]
		if ok && v != data {
			return false
		} else {
			tmp[k] = v
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
