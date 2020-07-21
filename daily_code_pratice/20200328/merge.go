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
// 要在原 nums1 上面操作,那么就才用i i+1 i+2 相互赋值的方法

// 一开始自己疯狂在试切片的插入，还整了threePoint.go arrIndexInsert.go　来测试
// 还在探查中
// 后来又想了先能解出测试用例，就用sort之后，将多余的0删除，发现哪里还是不对

//　看看评论和解题吧

import (
	"fmt"
	"sort"
)

func mergeComment(nums1 []int, m int, nums2 []int, n int) {

	nums1 = append(nums1[:m], nums2[:n]...)

	sort.Ints(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums1 = append(nums1, nums2...)
	fmt.Println(nums1)

	sort.Ints(nums1)
	fmt.Println("sort", nums1)

	zeros := len(nums1) - n

	fmt.Println(nums1[zeros:len(nums1)])

}

func insertStringSliceCopy(slice, insertion []int, index int) []int {
	result := make([]int, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &m, &n)
	var n1 []int
	var n2 []int

	fmt.Printf("n1:")
	for i := 0; i < m+n; i++ {
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

// 收获：[:...] 这个操作还是需要多熟悉啊，穿上点点就不认识了
