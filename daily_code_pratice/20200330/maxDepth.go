package main

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//       / \
//         9  20
//	     /  \
//	        15   7
//		返回它的最大深度 3 。
//

//个人思路
// 反正到树这块基本都是递归了，就继续参考评论和题解练习代码

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := 1, 1

	if root.Left != nil {
		l += maxDepth(root.Left)
	}

	if root.Right != nil {
		r += maxDepth(root.Right)
	}

	return maxNums(l, r)
}

func maxNums(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println("vim-go")
}
