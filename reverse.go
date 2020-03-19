package main

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
// 输入: -123
// 输出: -321
// 示例 3:
//
// 输入: 120
// 输出: 21
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。    2^31 =  2147483648

import (
	"fmt"
	"strconv"
)

//个人思路：
// 字符串分割
// 两个数组
// 字符串转整形，且需判断溢出
// go语言 Int 最大 num = 9223372036854775807

func reverse(x int) int {
	str_num := strconv.Itoa(x)

	var strs []string

	length := len(str_num)

	for i := length; i >= 0; i++ {
		if str_num[i] == '-' {
			continue
		}
		strs = append(strs, str_num[i])
	}

	res, err := strconv

}

func main() {
	var num int
	fmt.Printf("nums:")
	fmt.Scanf("%d", &num)

	res := reverse(num)
	fmt.Printf("reverse:")
	fmt.Println(res)
}
