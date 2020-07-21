package main

//给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//
//找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//
//您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//
//示例:
//
//输入:
//[4,3,2,7,8,2,3,1]
//
//输出:
//[5,6]

// 个人思路
// 排序，遍历，下标，下标+1 与当前下标对比

import (
	"fmt"
)

// by comment 马永真
func findDisappearedNumbers(nums []int) []int {
	var res []int

	index := 0
	for index < len(nums) {
		for nums[index] != index+1 && nums[index] != nums[nums[index]-1] {
			nums[index], nums[nums[index]-1] = nums[nums[index]-1], nums[index]
		}
		index++
	}

	for i, v := range nums {
		if v != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}

// 收获：先把元素移至自己所在下标，但出现两次的，什么情况呢
