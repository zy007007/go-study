package main

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。
//
//
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//
//示例:
//
//输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]

//个人思路
//要慢慢开始想较快的思路，考虑时间和内存
// 要在原 nums1 上面操作

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

}

func main() {
	var n, m int
	var n1, n2 []int
	fmt.Scanf("%d %d", &m, &n)

	fmt.Printf("n1:")
	for i := 0; i < m; i++ {
		var data int
		fmt.Scanf("%d", &data)
		n1 = append(n1, data)
	}

	fmt.Printf("n2:")
	for i := 0; i < n; i++ {
		var data int
		fmt.Scanf("%d", &data)
		n2 = append(n2, data)
	}

	merge(n1, m, n2, n)
	fmt.Println(n1)
}
