package main

//给定两个二叉树，编写一个函数来检验它们是否相同。

//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

// 个人思路
// 我看看啊，连续７题都是树相关
// 这个真认了，７题就对着评论和题解，这里练练代码感觉吧
// 整不动花里胡哨了

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func main() {
	fmt.Println("vim-go")
}

// 收获：注意递归退出的条件
