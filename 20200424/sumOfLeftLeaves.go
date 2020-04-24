package main

//计算给定二叉树的所有左叶子之和。
//
//示例：
//
//    3
//       / \
//         9  20
//	     /  \
//	        15   7
//
//		在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

// 个人思路:
// 如何判断一个叶子节点是左叶子

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right) + res

}

func main() {
	fmt.Println("vim-go")
}

// 收获：会了没，判断是否为左叶子
