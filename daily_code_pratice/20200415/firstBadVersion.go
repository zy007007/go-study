package main

//你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
//
//假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
//你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
//
//示例:
//
//给定 n = 5，并且 version = 4 是第一个错误的版本。
//
//调用 isBadVersion(3) -> false
//调用 isBadVersion(5) -> true
//调用 isBadVersion(4) -> true
//
//所以，4 是第一个错误的版本。

// 个人思路
// 翻译，判断1-n中，使用isBadVersion判断第一个false

import (
	"fmt"
)

// by official
func firstBadVersionV1(n int) int {
	for i := 0; i < n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return n
}

func firstBadVersionV2(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2 //  (l + r) / 2 的变形
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println("vim-go")
}

// 收获：理解很重要，学习新的二分查找的表达式
