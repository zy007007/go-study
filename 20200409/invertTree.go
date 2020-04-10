package main

//翻转一棵二叉树。
//
//示例：
//
//输入：
//
//     4
//        /   \
//	  2     7
//	   / \   / \
//	   1   3 6   9
//	   输出：
//
//	        4
//		   /   \
//		     7     2
//		      / \   / \
//		      9   6 3   1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeComment(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeComment(root.Left)
	invertTreeComment(root.Right)
	return root
}

// 收获：模板流的话，我这次可是记住了
