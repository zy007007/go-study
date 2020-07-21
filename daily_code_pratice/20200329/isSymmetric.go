package main

//给定一个二叉树，检查它是否是镜像对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//       / \
//         2   2
//	  / \ / \
//	  3  4 4  3
//	  但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//	      1
//	         / \
//		   2   2
//		      \   \
//		         3    3

//个人思路
// 又是要递归了嘛

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSame(root, root)
}

func isSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil || q != nil {
		return false
	}

	if p.Val == q.Val {
		return isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
	}
	return false
}

func main() {
}

// 收获：还有误，下周开始继续看看吧
