package main

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true

type ListNode struct {
	Val  int
	Next *ListNode
}

// by ElliotXX 思路：先使用快慢指针定位链表中点，然后反转中点的后半部分，最后分别从开头和中点处遍历比较是否是回文链表
func isPalindromeComment(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	pre, cur := nil, slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	mid := pre
	for mid != nil {
		if head.Val != mid.Val {
			return false
		}
		mid = mid.Next
		head = head.Next
	}
	return true
}

// 收获：学到了，指针快慢找中点
