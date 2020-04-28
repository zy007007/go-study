package main

//写一个程序，输出从 1 到 n 数字的字符串表示。
//
//1. 如果 n 是3的倍数，输出“Fizz”；
//
//2. 如果 n 是5的倍数，输出“Buzz”；
//
//3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
//
//示例：
//
//n = 15,
//
//返回:
//[
//    "1",
//        "2",
//	    "Fizz",
//	        "4",
//		    "Buzz",
//		        "Fizz",
//			    "7",
//			        "8",
//				    "Fizz",
//				        "Buzz",
//					    "11",
//					        "Fizz",
//						    "13",
//						        "14",
//							    "FizzBuzz"
//							    ]

// 个人思路
// 简单题，暴力解，%3 %5 = 0
// 数字转字符串输出

//　没仔细看题，还有双倍输出，就这样子吧
// 不过没想通为什么下面这样子写代码编译不通过，tmp 会报错
// 提交时还是把tmp定义在if前面了

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	var res []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			tmp := "Fizz"
		} else if i%5 == 0 {
			tmp := "Buzz"
		} else {
			tmp := strconv.Itoa(i)
		}
		res = append(res, tmp)
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
