// https://leetcode.cn/problems/merge-two-sorted-lists/
package solutions

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 针对空链表进行控制，如果遇到直接返回另一个链表头结点
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := list1
	if list2.Val < list1.Val {
		head = list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	}
	if list2 == nil {
		tail.Next = list1
	}
	return head
}

func TestMergeList(t *testing.T) {
	fmt.Println(mergeTwoLists(nil, nil))
}
