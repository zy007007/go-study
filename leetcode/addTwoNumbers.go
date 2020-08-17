/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}


func ListToNumber(l *ListNode) int64 {
	var num, res string

	for l != nil {
		num += string(l.Val)
		l = l.Next
	}

	length := len(num)
	for i=length-1; i>=0; i-- {
		res += num[i]
	}

	r, _ := strconv.ParseInt(res, 10, 64)
	return r
}


// 链表插入不会写
func NumberToList(n int64) *ListNode {
	str := strconv.FormatInt(n, 10)

	var head = new(ListNode)


}