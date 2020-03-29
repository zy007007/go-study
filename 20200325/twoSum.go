package main

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

import (
	"fmt"
	"sort"
)

//个人思路：
//对数组进行排序
//删除比target大的数
//从头(尾)开始，从尾(头)-1匹配相加
//若有结果，记录下标；无结果返回空

//按以上思路完成代码测试问题：输出的下标，是排序后的，而非原生数组对应的

//改进：
//符合条件的两元素，返回原数组下标，因此需要开始时有两个相同的数组，一个找数，一个找符合条件的下标

// 两组测试用例，符合预期结果
//arr len:5
//arr nums:12 1 3 4 5
//target nums:15
//0 2

//arr len:5
//arr nums:1 2 3 4 5
//target nums:3
//0 1

//错误用例
//[3,3]
//6
//输出
//[0,0]
//预期结果
//[0,1]

//待改进
func twoSum(nums []int, target int) []int {
	length := len(nums)

	copy_arr := make([]int, length)

	copy(copy_arr, nums)

	sort.Ints(nums) // 升序排序

	for n := length - 1; n >= 0; n-- {
		if nums[n] >= target {
			nums = nums[:len(nums)-1]
		}
	}

	var ok []int
	new_length := len(nums)
	optimization := 0

	for right := new_length - 1; right >= 0; right-- {
		for left := 0; left < new_length-optimization; left++ {
			if nums[right]+nums[left] == target {
				ok = append(ok, nums[left])
				ok = append(ok, nums[right])
				break
			}
		}
		optimization += 1
		if len(ok) != 0 {
			break
		}
	}

	var res []int
	if len(ok) == 2 {
		for v, k := range copy_arr {
			if k == ok[0] {
				res = append(res, v)

			} else if k == ok[1] {
				res = append(res, v)
			}

			if len(res) == 2 {
				break
			}

		}
	}

	return res
}

func main() {
	var l int
	var arrs []int

	fmt.Printf("arr len:")
	fmt.Scanf("%d", &l)

	fmt.Printf("arr nums:")
	for i := 0; i < l; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arrs = append(arrs, num)
	}

	var target int
	fmt.Printf("target nums:")
	fmt.Scanf("%d", &target)

	res := twoSum(arrs, target)

	for _, v := range res {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
