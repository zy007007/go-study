package main
//给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
//
//说明：不要使用任何内置的库函数，如  sqrt。
//
//示例 1：
//
//输入：16
//输出：True
//示例 2：
//
//输入：14
//输出：False

// 个人思路
// 又是数学题
// 这次我知道了，FormatInt(num, 2)
// 转为二进制，然后判断第一位为1，其余为0

// 不是老哥，我在想什么呢

import "fmt"

// by ly爱编程:1 , 4 = 1 + 3 , 9 = 1 + 3 + 5.......每次减去一个奇数，最后若为0 则为true否则为false
func isPerfectSquare(num int) bool {
	tmp := 1
	for num > 0 {
		num = num - tmp
		tmp += 2
	}

	if num == 0 {
		return true
	}
	return false
	
}


func main() {
	fmt.Println("vim-go")
}
//收获：学到了一个方法啊
