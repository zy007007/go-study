package main

//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//       / \
//         9  20
//	     /  \
//	        15   7
//		返回它的最小深度  2.

// 个人思路
// 个人情绪缓好了，这个题有读了一下
// 其实没太懂示例为什么是２，可能是对树的相关知识点不熟，并且题目描述不清

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min = 1<<63 - 1
	dfs(root, 1)
	return min
}

func dfs(node *TreeNode, depth int) {
	if node.Right == nil && node.Left == nil {
		if depth < min {
			min = depth
		}
		return
	}

	if node.Left != nil {
		dfs(node.Left, depth+1)
	}

	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}
