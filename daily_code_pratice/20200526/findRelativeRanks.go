package main

//给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold Medal", "Silver Medal", "Bronze Medal"）。
//
//(注：分数越高的选手，排名越靠前。)
//
//示例 1:
//
//输入: [5, 4, 3, 2, 1]
//输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
//解释: 前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal" and "Bronze Medal").
//余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。
//提示:
//
//N 是一个正整数并且不会超过 10000。
//所有运动员的成绩都不相同。

import (
	"fmt"
)

func findRelativeRanks(nums []int) []string {
	max1, max2, max3 := nums[0], nums[1], nums[2]
	for i := 3; i < len(nums-2); i++ {
		if max1 < nums[i] {
			max1 = nums[i]
		}

		if max2 < nums[i+1] {
			max2 = nums[i+1]
		}

		if max3 < nums[i+2] {
			max3 = nums[i+2]
		}
	}

	//
	sort.Int(nums)
	nums[0] = "Gold Medal"
	nums[1] = "Silver Medal"
	nums[2] = "Bronze Medal"
}

func main() {
	fmt.Println("vim-go")
}
