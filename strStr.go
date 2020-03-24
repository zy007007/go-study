package main

//实现 strStr() 函数。
//
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//
//示例 1:
//
//输入: haystack = "hello", needle = "ll"
//输出: 2
//示例 2:
//
//输入: haystack = "aaaaa", needle = "bba"
//输出: -1
//说明:
//
//当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
//对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

//个人思路
//就是经典的子串字符串位置
//可以继续采用 arr[i:j] ?= needle

//提交测试用例通过

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	res := -1
	needleLen := len(needle)
	for i := 0; i < len(haystack); i++ {
		if needleLen+i > len(haystack) {
			break
		}

		tmp := string(haystack[i:(needleLen + i)])
		if tmp == needle {
			res = i
			break
		}
	}
	return res
}

func main() {
	var longStr, shortStr string
	fmt.Printf("长字符串:")
	fmt.Scanf("%s", &longStr)

	fmt.Printf("子字符串:")
	fmt.Scanf("%s", &shortStr)

	fmt.Printf("输出:")
	fmt.Println(strStr(longStr, shortStr))
}

// 收获：临界值的多思考，希望下次类似情况可以一步考虑到位。另外，代码用了两个if，有点菜啊
