package main

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//示例 2:
//
//输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。

//个人思路
//初看很简单
//主要有如果是9，需要进位
// [9,9] [1,0,0]
// [1,9,9] [2,0,0]

// 从尾开始，如果是9，则进位，继续判断
// 结束标志是没有进位直接+1 退出，有的话，进一次，判断一次，如果一直进位到头，判断是不是9
// 用递归？

// 测试用例通过
// 看看评论

import (
	"fmt"
)

func plusOneComment(digits []int) []int {
	l := len(digits) - 1
	pls := false

	for i := l; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			pls = true
			continue
		}

		digits[i] = digits[i] + 1
		pls = false
		break
	}

	if pls {
		digits = append([]int{1}, digits...) // 这里学习了，头加?
	}

	return digits
}

func plusOne(digits []int) []int {

	l := len(digits) - 1
	digits[l] += 1
	for l >= 0 {
		if digits[l] == 10 {
			if l-1 < 0 {
				digits[l] = 0
				break
			}
			digits[l-1] += 1
			digits[l] = 0
		}
		l -= 1
	}
	var res []int
	if digits[0] == 0 {
		res = append(res, 1)
		for _, v := range digits {
			res = append(res, v)
		}

	} else {
		res = digits
	}

	return res
}

func main() {
	var length int
	var arr []int
	fmt.Printf("数组长度:")
	fmt.Scanf("%d", &length)

	fmt.Printf("输入元素:")
	for i := 0; i < length; i++ {
		var data int
		fmt.Scanf("%d", &data)

		arr = append(arr, data)
	}

	fmt.Printf("输出:")
	fmt.Println(plusOne(arr))

}

// 收获: 本来最后有修改怎么头加1，学习评论和提交记录代码，见line 54
