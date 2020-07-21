package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串

//个人思路：
//想了出现的规律：1.两两前后匹配 2.中心对称 3. 1+2
//然后发现要是让我开始写，估计就又是循环判断等等
//学习了评论中的一个方法，过一遍吧

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
		s = strings.Replace(s, "()", "", 1)
		s = strings.Replace(s, "{}", "", 1)
		s = strings.Replace(s, "[]", "", 1)
	}

	return s == ""
}

func main() {
	var str string
	fmt.Printf("输入：")
	fmt.Scanf("%s", &str)
	fmt.Printf("输出：")
	fmt.Println(isValid(str))
}

// 收获：注意功能函数是否有返回值
