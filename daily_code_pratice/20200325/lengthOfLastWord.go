package main

//给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
//
//如果不存在最后一个单词，请返回 0 。
//
//说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。
//
//
//
//示例:
//
//输入: "Hello World"
//输出: 5

//个人思路
//没懂滚动显示这个表述
//用内置函数的话，就按空格分割strings.Split() ?  数组len=1 返回　０　　否则返回数组-1元素　length,　　go 无-1

//不用内置，用遍历的话，那就从尾开始遍历，count 记录长度，遇到" "退出，返回　count，到头没有" "返回０

//还有高级点的吗

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lengthOfLastWord(s string) int {

	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			break
		}
		count += 1

	}

	if count == len(s) {
		count = 0
	}

	return count
}

func main() {
	//var s string
	fmt.Printf("输入:")
	//fmt.Scanf("%s", &s)
	input := bufio.NewReader(os.Stdin)

	s, _ := input.ReadString('\n')

	rs := strings.Trim(s, "\n")
	fmt.Printf("输出:")
	fmt.Println(lengthOfLastWord(rs))
}

//收获:输入的字符串包含空格，采用了新的方法，就是要来回删除换行，有待继续研究一下
