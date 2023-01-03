// https://leetcode.cn/problems/add-two-numbers/
package solutions

import (
	"fmt"
	"testing"
)

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
func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	fmt.Println(addTwoNumbers(l1, l2))
}
