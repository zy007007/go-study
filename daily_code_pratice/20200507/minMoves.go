package main

//给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动将会使 n - 1 个元素增加 1。
//
//示例:
//
//输入:
//[1,2,3]
//
//输出:
//3
//
//解释:
//只需要3次移动（注意每次移动会增加两个元素的值）：
//[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]

// 个人思路
// 找规律的题吧
// 先找本数组中最大的

import (
	"fmt"
)

// by comment 杨
// 刚开始碰上这题准备每进行一次循环进行一次插入排序，后面仔细思考是不是其他元素都加一的次数是不是等于本元素每次减一，直到所有的值到最小值呢。
func minMoves(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}

	sort.Ints(nums)
	count := 0
	for i := 1; i < l; i++ {
		count += nums[i] - nums[0]
	}
	return count
}

func main() {
	fmt.Println("vim-go")
}
