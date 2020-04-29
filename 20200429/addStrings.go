package main

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
//注意：
//
//num1 和num2 的长度都小于 5100.
//num1 和num2 都只包含数字 0-9.
//num1 和num2 都不包含任何前导零。
//你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

// 个人思路
// map 字符串 - 下标, 真实数字
// 万物皆可map

import (
	"fmt"
)

// by comment Super玛丽
func addStrings(num1 string, num2 string) string {
	nr1 := []rune{num1}
	nr2 := []rune{num2}
	result := make([]rune, 0)

	l1 := len(nr1) - 1
	l2 := len(nr2) - 1
	overflow := 0

	if l1 >= 0 || l2 >= 0 {
		v1 := 0
		v2 := 0
		if l1 >= 0 {
			v1 = int(nr1[l1] - '0')
		}

		if l2 >= 0 {
			v2 = int(nr2[l2] - '0')
		}

		value := v1 + v2 + overflow
		overflow = value / 10
		value = value % 10
		result = append([]rune{strconv.Itoa(value)}, result...)

		l1--
		l2--
	}

	if overflow > 0 {
		result = append([]rune{"1"}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println("vim-go")
}

// 收获：熟悉代码
