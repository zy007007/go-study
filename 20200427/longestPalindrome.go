package main

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
//在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
//注意:
//假设字符串的长度不会超过 1010。
//
//示例 1:
//
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

// 个人思路
// 字母出现次数>=2

// by comment Dylan
func longestPalindrome(s string) int {
	//letterArr := [26]byte{'a', 'b', 'c', 'd', 'e', ''}
	apl, res := [52]uint8{}, 0
	for i := range s {
		k := s[i] - 'a' + 'Z' - 'A' // s[i] 小写
		if s[i] <= 'Z' {
			k = s[i] - 'A' // s[i] 大写
		}

		if apl[k]++; apl[k]&1 == 0 { // 标志位，字母存在，>=二次出现
			res += 2
		}
	}

	if len(s)-res > 0 { // 是否存在单一中心值
		res++
	}
	return res

}

// 收获：uint8  ACSII码
