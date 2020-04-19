package main
//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4]
//说明:
//
//输出结果中的每个元素一定是唯一的。
//我们可以不考虑输出结果的顺序。

// 个人思路
// 交集，排序，遍历对比

import "fmt"


// by 青泥何盘盘 map
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range nums1 {
		set[v] = true
	}

	for _, v1 := range nums2 {
		if d, ok := set[v1]; ok && d{
			res = append(res, v1)
			set[v1] = false
		}
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}
