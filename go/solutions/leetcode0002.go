// https://leetcode.cn/problems/add-two-numbers/
package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	head = &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		cur.Next = &ListNode{}
		if cur.Val > 9 {
			cur.Next.Val = cur.Val / 10
			cur.Val = cur.Val % 10
		}
		cur = cur.Next
		if l1 != nil {
			cur.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cur.Val += l2.Val
			l2 = l2.Next
		}
	}
	if cur.Val > 9 {
		cur.Next = &ListNode{Val: cur.Val / 10}
		cur.Val = cur.Val % 10
	}
	return head.Next
}
