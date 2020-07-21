package main

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器，且 n 的值至少为 2。
//
//图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
//示例：
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49

// 1 7 - 1 * 7
//   1 3 - 1 * 6
//   8 7 - 7 * 7
// 选 8 7 - 7 * 7
//      6 7 -　6 * 6
//	8 3 - 3 * 3
// 选 8 7 - 7 * 7
//      2 7 - 2 * 2

// area = min(height[i], height[j]) * (j-j)

import (
	"fmt"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		h := min(height[i], height[j])
		res = max(res, h*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
