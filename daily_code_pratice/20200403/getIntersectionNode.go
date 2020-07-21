package main

//编写一个程序，找到两个单链表相交的起始节点。
//
//注意：
//
//如果两个链表没有交点，返回 null.
//在返回结果后，两个链表仍须保持原有的结构。
//可假定整个链表结构中没有循环。
//程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

// 个人思路
// 感觉是简单题了吧
// 遍历，判断值是否相等

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	for headA.Next != nil ||　headB.Next != nil {
		if headA.Val == headB.Val {
			return headA
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}
	return nil
}

func getIntersectionNodeComment(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func main() {
	fmt.Println("vim-go")
}
