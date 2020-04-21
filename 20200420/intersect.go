package main

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2,2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [4,9]
//说明：
//
//输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//我们可以不考虑输出结果的顺序。
//进阶:
//
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果 nums1 的大小比 nums2 小很多，哪种方法更优？
//如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

// 个人思路
// 次数一致，4 9 在nums2中，怎么叫次数一致呢
//

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	right1, left1, right2, left2 := 0, len(nums1), 0, len(nums2)
	var res []int
	for right1 < left1 && right2 < left2 {
		if nums1[right1] == nums2[right2] {
			res = append(res, nums1[right1])
			right1++
			right2++
		} else if nums1[right1] > nums2[right2] {
			right2++
		} else {
			right1++
		}
	}
	return data
}

func main() {
	fmt.Println("vim-go")
}

// 收获：正常的解题和代码思路，只是依然未理解什么叫次数一致
