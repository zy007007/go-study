package main

//给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
//
//这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
//
//示例1:
//
//输入: pattern = "abba", str = "dog cat cat dog"
//输出: true
//示例 2:
//
//输入:pattern = "abba", str = "dog cat cat fish"
//输出: false
//示例 3:
//
//输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false
//示例 4:
//
//输入: pattern = "abba", str = "dog dog dog dog"
//输出: false
//说明:
//你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

// 个人思路
// map1[a] =  1
// map2[dog] = 1
// map1[b] = 2
// map2[cat] = 2
// 判断　map1 和　map2 是否相等

// by comment ElliotXX
func wordPattern(pattern string, str string) bool {
	str_arr := strings.Fields(str)
	if len(pattern) != len(str_arr) {
		return false
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		v1, ok1 := m1[pattern[i]]
		v2, ok2 := m2[str_arr[i]]

		if ok1 && v1 != str_arr[i] || ok2 && v2 != pattern[i] {
			return false
		} else {
			m1[pattern[i]] = str_arr[i]
			m2[str_arr[i]] = pattern[i]
		}
	}
	return true
}
