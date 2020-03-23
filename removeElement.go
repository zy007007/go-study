package main

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//
//
//示例 1:
//
//给定 nums = [3,2,2,3], val = 3,
//
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//
//你不需要考虑数组中超出新长度后面的元素。
//示例 2:
//
//给定 nums = [0,1,2,2,3,0,4,2], val = 2,
//
//函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
//
//注意这五个元素可为任意顺序。
//
//你不需要考虑数组中超出新长度后面的元素。

//个人思路
//原地的话，那就遍历到val，把后面的依次前移
// 定位到要删除的元素下标k，用 [0:k] + [k+2:-1]  拼接方法，实现删除，但是算原地删除吗，这前后两个拼接的数组用的地址，如果是原数组的地址，那就没问题

// 发现可以用一个简单的办法，结果，测试用例过不去
// 我看题最后要的是结果长度啊，函数返回也是int类型

import (
	"fmt"
)

func removeElement(nums []int, val int) int {

	//k := 1
	count := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			//nums = nums[0:k-1] + nums[k:len(nums)]

			//nums = nums[0:k-1]
			//nums = append(nums, nums[k:len(nums)])
			count -= 1

		}
	}
	return count
}

func main() {
	var length, val int
	var arr []int
	fmt.Printf("数组长度:")
	fmt.Scanf("%d", &length)

	fmt.Printf("输入元素:")
	for i := 0; i < length; i++ {
		var data int
		fmt.Scanf("%d", &data)

		arr = append(arr, data)
	}

	fmt.Printf("移除元素:")
	fmt.Scanf("%d", &val)

	fmt.Printf("输出数组长度:")
	fmt.Println(removeElement(arr, val))
}

// 收获：go语言数组下标获取，目前看起来是没有[-1]这种的；一个数组，[x:y] 分成两个数组 ，地址用的是老的吧，(可以区验证一下，打印地址)，那么如果可以实现分成的两个数组直接合并，那么也算是原地处理
