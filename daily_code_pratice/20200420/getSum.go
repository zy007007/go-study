package main
//不使用运算符 + 和 - ，计算两整数 a 、b之和。
//
//示例 1:
//
//输入: a = 1, b = 2
//输出: 3
//示例 2:
//
//输入: a = -2, b = 3
//输出: 1

// 个人思路
// 应该知道用什么，但是想不到


// by comment 王小宁
func getSum(a int, b int) int {
	for b != 0 {
		tmp := a ^ b
		b = (a & b) << 1
		a = tmp
	}
	return a
}
// 收获：熟悉代码吧
