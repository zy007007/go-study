package main

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
//本题中，一棵高度平衡二叉树定义为：
//
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//
//示例 1:
//
//给定二叉树 [3,9,20,null,null,15,7]
//
//    3
//       / \
//         9  20
//	     /  \
//	        15   7
//		返回 true 。
//
//		示例 2:
//
//		给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//		       1
//		             / \
//			          2   2
//				      / \
//				         3   3
//					   / \
//					    4   4
//					    返回 false 。

// 个人思路
// 反正肯定又是递归

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return root == nil || isBalanced(root.Left) && math.Abs(height(root.Left)-height(root.Right)) < 2 && isBalanced(root.Right)
}

func height(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	return math.Max(height(root.Left), height(root.Right)) + 1
}
