package main

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

//个人思路
//这是经典的题了吧
//大风向就是动态规划，进一步下去就是斐波那契数列的第n项是多少

import (
	"fmt"
)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	i1, i2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := i1 + i2
		i1 = i2
		i2 = tmp
	}
	return i2
}

func main() {
	var n int
	fmt.Printf("输入:")
	fmt.Scanf("%d", &n)

	fmt.Printf("输出:")
	fmt.Println(climbStairs(n))
}

// 收获:到达第n阶的方法和，等于到 n-1 的方法和 加上 n-2 的方法和： res[n] = res[n-1] + res[n-2]
//res[1] = 1
//res[2] = 2
//for i:=3;i<=n;i++{
//	res[i] = res[i-1] + res[i-2]
//}
//return res[n]
