package main

//删除链表中等于给定值 val 的所有节点。
//
//示例:
//
//输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5

// 个人思路
// 是基础题
// 遍历，判断val相等否，然后next - nextnext

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	cur := head
	for cur != nil {
		if cur.Val == val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}

	return cur
}

func removeElements(head *ListNode, val int) *ListNode {
	tmp := new(ListNode)
	tmp.Next = head
	pre, cur := tmp, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = curr.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return tmp.Next
}

// 收获：tmp.Next 指向的是 pre  ，在遍历cur时，又是给pre加节点
