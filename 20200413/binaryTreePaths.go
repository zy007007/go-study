package main

//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//输入:
//
//   1
//    /   \
//    2     3
//     \
//       5
//
//       输出: ["1->2->5", "1->3"]
//
//       解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

// 个人思路
// 找到叶子节点-》一直找他的root，直到头部

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = []string{}
	if root == nil {
		return res
	}

	dfs(root, nil)
	return res
}

func dfs(root *TreeNode, path []string) {
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		res = append(res, strings.Join(path, "->"))
		return
	}

	if root.Left != nil {
		dfs(root.Left, path)
	}

	if root.Right != nil {
		dfs(root.Right, path)
	}
}
