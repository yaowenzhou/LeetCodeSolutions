// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	curNode := head.Next
	tmpHead := head
	tmpTail := head
	tmpCnt := 1
	for ; curNode != nil; curNode = curNode.Next {
		tmpTail.Next = curNode // curNode追加到临时链表
		tmpTail = tmpTail.Next
		if tmpCnt < n+1 { // 临时链表长度少于n+1，则只追加，并将临时链表长度+1
			tmpCnt++
		} else { // 临时链表长度>=n+1，则还需要弹出临时链表的头结点
			tmpHead = tmpHead.Next
		}
	}
	if tmpCnt < n { // 临时链表长度不足n，说明整个链表的长度都不足n,不用删
		return head
	}
	// 此时
	if tmpCnt == n {
		return head.Next
	}
	tmpHead.Next = tmpHead.Next.Next
	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	fmt.Println(removeNthFromEnd(head, 2))
}
