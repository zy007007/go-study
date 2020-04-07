package main
//编写一个算法来判断一个数是不是“快乐数”。
//
//一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。
//
//示例: 
//
//输入: 19
//输出: true
//解释: 
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1

// 个人思路
// 一开始想的是递归吧
// 看了评价和题解，根本不是用递归
// 

import (
	"fmt"
)

func isHappy(n int) bool {
	if n == 1 {
		return true
	} else {
		return false
	}	  

	return isHappy(n[0]^2 + n[1]^2 + ... + n[len-1]^2)
}


func isHappy(n int) bool {
	hash := make(map[int]bool)

	for n != 1 {
		if _, ok := hash[n]; ok {
			return false
		} 
		hash[n] = true
		next := 0
		for n != 0 {
			next += (n%10) * (n%10)
			n /= 10
		}
		n = next
	}
	return true
}


func main() {
	fmt.Println("vim-go")
}
