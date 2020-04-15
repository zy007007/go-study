package main
//给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
//
//示例 1:
//
//输入: [3,0,1]
//输出: 2
//示例 2:
//
//输入: [9,6,4,2,3,5,7,0,1]
//输出: 8
//说明:
//你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

// 个人思路
// 求和

import "fmt"


// by ElliotXX
func missingNumber(nums []int) int {
	length := len(nums)

	sum1, sum2 := length, 0

	for i:=0; i<length; i++ {
		sum1 += i
		sum2 += nums[i]
	}

	return sum1 - sum2
}


func main() {
	fmt.Println("vim-go")
}
// 收获：这个秒啊，一次循环两个sum
