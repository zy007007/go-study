package main

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 个人思路
// 逆置链表
// 输入的链表，从头插变尾插，或者反过来呗

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListComment(head *ListNode) *ListNode {
	var newHead, next *ListNode

	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}
