package main

//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
//示例 1:
//
//输入: nums = [1,2,3,1], k = 3
//输出: true
//示例 2:
//
//输入: nums = [1,0,1,1], k = 1
//输出: true
//示例 3:
//
//输入: nums = [1,2,3,1,2,3], k = 2
//输出: false

// 个人思路
// 暴力解法，问题应该不大
// 再想想其他
// 把一个数组，分成两个
// 哦我忽略了k

import (
	"fmt"
)

// by YL
func containsNearbyDuplicateComment(nums []int, k int) bool {
	var m = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			t := 0
			if m[nums[i]] > i {
				t = m[nums[i]] - i
			} else {
				t = i - m[nums[i]]
			}
			if t <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}

// 收获：map大法好
