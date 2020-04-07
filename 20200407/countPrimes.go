package main

//统计所有小于非负整数 n 的质数的数量。
//
//示例:
//
//输入: 10
//输出: 4
//解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

// 个人思路
// 基础的写了

func countPrimes(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if isPrim(i) {
			count++
		}
	}
	return count
}

func isPrim(n int) bool {
	res := true

	tmp := math.Sqrt(n)

	for i := 2; i < tmp; i++ {
		if n%i == 0 {
			res = false
			break
		}
	}
	return res
}
