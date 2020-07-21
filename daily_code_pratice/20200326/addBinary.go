package main

//给定两个二进制字符串，返回他们的和（用二进制表示）。
//
//输入为非空字符串且只包含数字 1 和 0。
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"

//个人思路
//转换为10进制，相加后，再转换为2进制。还能顺便写两个2,10进制转换的函数，万一以后还能用呢

//测试用例，下面这个出错
// a "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
// b "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
// 私以为是math.Pow(2) 的问题

// 看评论和题解
// 大风向是把短的先补齐，具体实现待看

import (
	"fmt"
	"math"
	"strconv"
)

func addBinary(a string, b string) string {

	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}

	A := twoToTen(a)
	B := twoToTen(b)

	return tenToTwo(A + B)
}

func twoToTen(n string) int {

	l := len(n)
	k := l - 1
	sum := 0
	for i := 0; i < l; i++ {
		if string(n[i]) == "1" {
			tmp := math.Pow(2, float64(k))

			sum += int(tmp)
		}
		k -= 1
	}
	return sum
}

func tenToTwo(n int) string {
	var tmp string

	for n != 0 {
		last := n % 2

		tmp += strconv.Itoa(last)
		n = n / 2
	}

	var res string

	for i := len(tmp) - 1; i >= 0; i-- {
		res += string(tmp[i])
	}

	return res
}

func main() {
	var a, b string
	fmt.Printf("输入二进制a:")
	fmt.Scanf("%s", &a)

	fmt.Printf("输入二进制b:")
	fmt.Scanf("%s", &b)

	fmt.Printf("输出:")
	fmt.Println(addBinary(a, b))
}

//收获：倒也复习了二十进制间的转换
