package main

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

// 个人思路
// 原数组操作，那就是遇到0了，前后前后比较的，移动到最后
// 或者找到０了，记录下标位置

import (
	"fmt"
)

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] == 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}

	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

func main() {
	fmt.Println("vim-go")
}
