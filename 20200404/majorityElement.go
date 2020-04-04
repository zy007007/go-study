package main

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2

// 个人思路
// 大的思路看，概率问题
// n 出现的次数，占长度的比例　>= 1/2
// 小的思路看，编程如何求概率
// 依次循环判断，每个数出现的次数，统计过得置空
// 结果存入map，结果类型怎么存，float
// 感觉能写出来，就是没用比较高级的算法耶

import "fmt"

func majorityElement(nums []int) int {
	var res int
	var str []string
	statitics := make(map[string]float64)

	str = nums // 将整形切片，转为字符切片

	length := len(nums)
	i := 0
	count := 0

	for i := 0; i < length; i++ {
		init := str[i]
		for j := i + 1; j < length; j++ {
			if str[j] == "x" {
				continue
			}

			if init == str[j] {
				count++
				str = "x"
			}
		}
		statitics[str[i]] = count / length
		count = 0
	}

	for v, k := range statitics {
		if k >= 0.5 {
			res = int(v)
			break
		}
	}

	return res
}

// 看了评论，果然有更好的方法
// 设一个标示已经标示位对应的次数
func majorityElement(nums []int) int {
	res := 0
	count := 0

	for _, m := range nums {
		if count == 0 {
			res = m
		}

		if m == res {
			count++
		} else {
			count--
		}

	}
	return res
}

func main() {
	fmt.Println("vim-go")
}

// 收获：方法肯定有，看你想的到吗
