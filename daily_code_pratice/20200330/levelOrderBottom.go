package main

//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//       / \
//         9  20
//	     /  \
//	        15   7
//		返回其自底向上的层次遍历为：
//
//		[
//		  [15,7],
//		    [9,20],
//		      [3]
//		      ]

//继续参考题解了，熟悉编写代码

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := list.New()

	queue.PushFront(root) // 头插

	for queue.Len() > 0 {
		var currentLevel []int
		queueLen := queue.Len()

		for i := 0; i < queueLen; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)

			currentLevel = append(currentLevel, node.Val) // 尾部移除的元素，放入当前层级的数组

			if node.Left != nil {
				queue.PushFront(node.Left)
			}

			if node.Right != nil {
				queue.PushFront(node.Right)
			}

		}
		res = append([][]int{currentLevel}, res...)
	}
	return res
}

// 收获：list 这个包，有双向队列的一些方法，https://golang.org/pkg/container/list/，头插尾插，头出尾出等等
