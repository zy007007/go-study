package main

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//     [1],
//         [1,1],
//	    [1,2,1],
//	      [1,3,3,1],
//	       [1,4,6,4,1]
//	       ]

//个人思路

//找规律吗感觉
//还是按着题解一个答案写的
//对我而言关键在于前几行，数组的相关定义

import "fmt"

func generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	for i := range arr {
		arr[i] = make([]int, i+1)
	}

	if numRows == 0 {
		return arr
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i-1][j-1]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println("vim-go")
}

// make 用来创建切片，map 和channel,所以这里一看是make，那么就很肯定是创建切片
