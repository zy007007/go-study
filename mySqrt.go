package main

//实现 int sqrt(int x) 函数。
//
//计算并返回 x 的平方根，其中 x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2
//示例 2:
//
//输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。

//个人思路
//求n平方根啊
// 从 1-n   开始遍历i，如果 i^2  < n   i++  ，i^2 > n , break   return i-=1

// 提交测试，结果超出时间限制
// 自己测试正常，没什么问题

// 看评论和题解，核心考点有二分法

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	var res int

	for i := 1; i <= x; i++ {

		if math.Pow(float64(i), float64(2)) == float64(x) {
			res = i
		} else if math.Pow(float64(i), float64(2)) > float64(x) {
			res = i - 1
		} else {
			continue
		}

		if res != 0 {
			break
		}
	}

	return res
}

func main() {
	var n int
	fmt.Printf("输入:")
	fmt.Scanf("%d", &n)

	fmt.Printf("输出:")
	fmt.Println(mySqrt(n))
}

// 收获:偏数学类型了，俗话说的好，生活处处皆数学，说起数学，就想到了高中的数学老师
