package main

//给定一个整数 n，返回 n! 结果尾数中零的数量。
//
//示例 1:
//
//输入: 3
//输出: 0
//解释: 3! = 6, 尾数中没有零。
//示例 2:
//
//输入: 5
//输出: 1
//解释: 5! = 120, 尾数中有 1 个零.
//说明: 你算法的时间复杂度应为 O(log n) 。

// 个人思路
// log n
// 如何判断0 的个数
// 一次遍历求阶乘，时间复杂度是　O(n) > O(log n)吧
// 所以需要边遍历边算结果
// 也不对，那还是O(n) 吧
// 所以还有更巧的方法

// 看评论大风向，求因子５的个数哟

func trailingZeroes(n int) int {
	var res int

	init := 1
	for i := n; i >= 1; i++ {
		tmp := i * (i - 1)
		if string(tmp)[len(tmp):][0] == "0" {
			res += 1
		}
	}

	return res
}

func trailingZeroesComment(n int) int {
	if n < 5 {
		return 0
	}

	return n/5 + trailingZeroesComment(n/5)
}
