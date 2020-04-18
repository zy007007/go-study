package main

//给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
//
//示例：
//
//给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
//
//sumRange(0, 2) -> 1
//sumRange(2, 5) -> -1
//sumRange(0, 5) -> -3
//说明:
//
//你可以假设数组不可变。
//会多次调用 sumRange 方法。

// 个人思路
// 设计数据结构问题，要我我就是常规

import "fmt"

// by comment ChuckieZhu
type NumArray struct {
	Value []int
}

func Constructor(nums []int) NumArray {
	num1 := NumArray{[]int{0}}
	for i, v := range nums {
		num1.Value = append(nums1.Value, v+num1.Value[i])
	}
	return num1
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Value[j+1] - this.Value[i]
}

func main() {
	fmt.Println("vim-go")
}

// 收获：结构体初始化；求数组下标和，题解的这个方法没有用常规的截取[i:j]
