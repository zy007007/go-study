package main

//给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
//
//示例 1:
//
//输入: 1
//输出: true
//解释: 20 = 1
//示例 2:
//
//输入: 16
//输出: true
//解释: 24 = 16
//示例 3:
//
//输入: 218
//输出: false

// 个人思路
// /2 /2 /2 /2 10进制换成二进制，判断是否全为1

import (
	"fmt"
)

// by 蝎子 n如果是2的幂, 二进制下最高位为1, 其余全是0; n-1要么等于0, 要么二进制下都是1; n&(n-1)必为0
func isPowerOfTwoComment(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println("vim-go")
}
