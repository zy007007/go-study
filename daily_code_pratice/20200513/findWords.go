package main
//给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
//示例：
//
//输入: ["Hello", "Alaska", "Dad", "Peace"]
//输出: ["Alaska", "Dad"]
//
// 个人思路
// 简单方法：
// 这，每一行先定义好
// 遍历单个字符串，判断
//
// 规律：

import "fmt"

// by comment 天驰
func findWords(words []string) []string {
	indexes := [26]int{}
	for _, v := range "qwertyuiop" {
		indexes[v-'a'] = 1
	}

	for _, v := range "asdfghjkl" {
		indexes[v-'a'] = 2
	}

	for _, v := range "zxcvbnm" {
		indexes[v-'a'] = 3
	}

	res := make([]string, 0)
	for _, word := range words {
		tmp := strings.ToLower(word)
		found := true
		for _, c := range tmp {
			if indexes[c-'a'] != indexes[tmp[0]-'a'] {
				found = false
				break
			}
		}
		if found {
			res = append(res, word)
		}
	}
	return res
}


func main() {
	fmt.Println("vim-go")
}
// 收获：以每个单词首字母为标志位，判断对应的index一直是首字母对齐吗，和首字母都是1，那就是第一行，首字母都是2，那就是第二行，首字母都是3，那就是第三行
