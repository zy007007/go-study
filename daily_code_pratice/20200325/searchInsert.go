package main

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。
//
//示例 1:
//
//输入: [1,3,5,6], 5
//输出: 2
//示例 2:
//
//输入: [1,3,5,6], 2
//输出: 1
//示例 3:
//
//输入: [1,3,5,6], 7
//输出: 4
//示例 4:
//
//输入: [1,3,5,6], 0
//输出: 0

//个人思路:
//给定的是排序好的数组
//无重复元素
//先判断是否==
//再判断是否 < arr[x]  是的话，插入
// 开始写发现两个可以合并

// 测试用例通过，再看评论发现，其实需要考虑数据量太大的时候，这题其实考的是二分法
import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			res = i
			break
		}
	}
	return res
}

func main() {
	var l, val int
	var arr []int
	fmt.Printf("输入数组长度：")
	fmt.Scanf("%d", &l)

	fmt.Printf("排序数组：")
	for i := 0; i < l; i++ {
		var data int
		fmt.Scanf("%d", &data)
		arr = append(arr, data)
	}

	fmt.Printf("目标值：")
	fmt.Scanf("%d", &val)

	fmt.Printf("输出：")
	fmt.Println(searchInsert(arr, val))
}

// 收获：考虑大数据量，二分法可以提高效率和时间下
