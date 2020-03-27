package main

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3

//个人思路
//不是，这个题又要先输入链表了，感觉可以整合到之前那个链表代码中去算了
//移步至 mergeTwoLists.go
//但是函数模块的代码，还是参考评论和题解，写写吧

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}
	return head
}

func main() {
	fmt.Println("vim-go")
}
