package main

//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
//示例:
//
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//           / \
//	      -3   9
//	         /   /
//		  -10  5

// 个人思路
// 熟悉写代码的感觉吧

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return &TreeNode{nums[len(nums)/2], sortedArrayToBST(nums[:len(nums)/2]), sortedArrayToBST(nums[len(nums)/2+1:])}
}

func main() {
	fmt.Println("vim-go")
}
