package main

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

//个人思路
//前段时间有看过，用了递归
// 定义一个新的链表
// 判断两个链表第一个元素的大小
// 加入新的链表中，然后递归进行下两个目标的比较

import (
	"fmt"
)

type ListNode struct {
	Data int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode

	// 若l1为空，返回l2，反之亦然
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	//　使用递归
	if l1.Data > l2.Data {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}

	return res
}

// 头插法
func addListData(nums int) *ListNode {
	var head = new(ListNode)
	address := head // address 记录头结点的地址

	for i := 0; i < nums; i++ {
		var n int
		fmt.Scanf("%d", &n)
		node := ListNode{Data: n}

		node.Next = address
		address = &node

	}

	return address
}

// 尾插法
func addListTailData(nums int) *ListNode {
	var head *ListNode

	tail := head

	for i := 0; i < nums; i++ {
		var m int
		fmt.Scanf("%d", &m)
		node := ListNode{Data: m}
		fmt.Println(tail)

		tail = &node
	}

	return tail
}

func showList(l *ListNode) {
	for l.Next != nil { // 如果不加.Next,输出的结果则包含空的头节点
		fmt.Println(l.Data, l.Next)
		l = l.Next
	}
}

func main() {
	var l1, l2 int
	fmt.Printf("l1 length:")
	fmt.Scanf("%d", &l1)
	fmt.Printf("l1 datas:")
	list1 := addListTailData(l1)
	showList(list1)

	fmt.Printf("l2 length:")
	fmt.Scanf("%d", &l2)
	fmt.Printf("l2 datas:")
	list2 := addListData(l2)
	showList(list2)

	//res := mergeTwoLists(list1, list2)
	//showList(res)
}

// 收获：熟悉了链表的定义，遍历，关键还是头插法和尾插法，怎么定义貌似决定了很大程度的代码编写和理解，见参考testList.go
