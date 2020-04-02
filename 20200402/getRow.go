package main

//给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 3
//输出: [1,3,3,1]

// 个人思路
// 输入的k是索引
// 转化成初代题，就是输入的 n+1 行的杨辉三角
// 结果输出最后一行即可
// 先自己尝试写昨天晚上练习的杨辉三角初代

// 自己基本可以写出来了，就是初始化切片时，需要去掉 []int{} 的 {}

import (
	"fmt"
)

func getRow(rowIndex int) []int {

	arr := make([][]int, rowIndex+1)

	for i := range arr {
		arr[i] = make([]int, i+1)
	}

	for x := 0; x < rowIndex+1; x++ {
		for y := 0; y <= x; y++ {
			if y == 0 || y == x {
				arr[x][y] = 1
			} else {
				arr[x][y] = arr[x-1][y] + arr[x-1][y-1]
			}
		}
	}

	return arr[rowIndex]

}

func main() {
	fmt.Println("vim-go")
}

// 收获：昨晚参考题解练习了，今天倒是可以基本顺利写出来，最近状态确实不错，保持吧
