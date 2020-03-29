package main

//给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
//
//
//示例 1:
//
//给定数组 nums = [1,1,2],
//
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//
//你不需要考虑数组中超出新长度后面的元素。
//示例 2:
//
//给定 nums = [0,0,1,1,1,2,2,3,3,4],
//
//函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//
//你不需要考虑数组中超出新长度后面的元素。

//个人思路
//原地删除啊
// 即可想到的有存入map,返回map长度
// 借助map测试成功，但是不符合题意

import "fmt"

func removeDuplicates(nums []int) int {
	dict := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := dict[nums[i]]; ok {
			continue
		} else {
			dict[nums[i]] = 1
		}
	}

	return len(dict)
}

func main() {
	var n int
	fmt.Printf("数组长度:")
	fmt.Scanf("%d", &n)

	var arr []int
	fmt.Printf("输入元素:")
	for i := 0; i < n; i++ {
		var data int
		fmt.Scanf("%d", &data)
		arr = append(arr, data)
	}

	fmt.Printf("输出")
	fmt.Println(removeDuplicates(arr))
}

//收获：熟悉map的使用
