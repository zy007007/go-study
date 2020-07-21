package main

//给你一个由若干 0 和 1 组成的数组 nums 以及整数 k。如果所有 1 都至少相隔 k 个元素，则返回 True ；否则，返回 False

//个人思路
// 单次循环，首位0 or 1

func kLengthApart(nums []int, k int) bool {

	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] != num[i+1] {
			count++
		}
		if k < count {
			return false
		}
	}
	return true
}
