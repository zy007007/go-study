package main

//给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
//
//假定 BST 有如下定义：
//
//结点左子树中所含结点的值小于等于当前结点的值
//结点右子树中所含结点的值大于等于当前结点的值
//左子树和右子树都是二叉搜索树
//例如：
//给定 BST [1,null,2,2],
//
//   1
//       \
//            2
//	        /
//		   2
//		   返回[2].
//
//		   提示：如果众数超过1个，不需考虑输出顺序
//

import (
	"fmt"
)

func findModeMyself(root *TreeNode) []int {
	count := make(map[int]int)
	var res []int
	if root == nil {
		return res
	}
}

func findMode(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	var max int
	m := make(map[int]int)
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		m[node.Val]++
		if max < m[node.Val] {
			max = m[node.Val]
			res = nil
			res = append(res, node.Val)
		} else if max == m[node.Val] {
			res = append(res, node.Val)
		}

		root = node.Right
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
