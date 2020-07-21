package main

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//示例 1:
//
//输入: 121
//输出: true
//示例 2:
//
//输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//示例 3:
//
//输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。

//个人思路
//整形转字符串
//从头，从尾开始比较
//注意比较的元素类型

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str_num := strconv.Itoa(x)

	left := 0
	right := len(str_num) - 1
	res := true

	for left <= right {
		if string(str_num[left]) != string(str_num[right]) {
			return false
			break
		}
		left += 1
		right -= 1
	}
	return res
}

func main() {
	var num int
	fmt.Printf("输入:")
	fmt.Scanf("%d", &num)
	fmt.Printf("输出:")
	fmt.Println(isPalindrome(num))
}
