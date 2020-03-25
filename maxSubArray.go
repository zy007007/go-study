package main

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

//个人思路
//低阶代码貌似要实现是简单
// 子数组从 k=1 开始，递增到length(arr)
// 将原数组从 k 开始的子数组，求和，放到新的数组中
// 最后对新的数组进行取最大值
//    新创了数组
//    时间复杂度是n*n吧

//然后，看了评论，动态规划
// 1.有个递推关系  fn = max(f(n-1)+fn,  fn)
// 用边遍历边求最大值的方法

import (
	"fmt"
)

func maxSubArray(nums []int) int {

	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= 0 {
			nums[i] += nums[i-1]
		}
	}

	max_num := nums[0]
	for j := 1; j < len(nums); j++ {
		if max_num < nums[j] {
			max_num = nums[j]
		}
	}

	return max_num
}

func main() {
	var n int
	fmt.Printf("数组长度:")
	fmt.Scanf("%d", &n)

	fmt.Printf("元素:")
	var arr []int
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		arr = append(arr, tmp)
	}

	fmt.Printf("输出:")
	fmt.Println(maxSubArray(arr))
}

// 收获：开始多方法学习
