// https://leetcode.cn/problems/merge-two-sorted-lists/
package solutions

import (
	"fmt"
	"leetcode_solutions_go/solutions/common"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	// 针对空链表进行控制，如果遇到直接返回另一个链表头结点
	if list1 == nil { // list为空直接返回list2即可
		return list2
	}
	if list2 == nil { // list2为空，直接返回list1即可
		return list1
	}
	head := list1              // 假定list1的头结点为待返回数据的头结点
	if list2.Val < list1.Val { // 如果list2的头结点数据更小
		head = list2       // 则将list2的头结点作合并链表的头结点
		list2 = list2.Next // 更新list2的头结点
	} else { // 合并链表的头结点依然是list1的头结点，需要更新list1的头结点
		list1 = list1.Next
	}
	tail := head                       // 尾结点
	for list1 != nil && list2 != nil { // 同时遍历list1和list2
		if list1.Val <= list2.Val { // 将二者中较小的头结点追加到尾结点，并更新该剩余链表的头结点
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next // 更新尾结点
	}
	// 如果list1为空，则将list2剩余部分追加到返回链表
	if list1 == nil {
		tail.Next = list2
	}
	// 如果list2为空，则将list1剩余部分追加到返回链表
	// 如果list1和list2同时为空也无所谓，无非是tail.Next 由nil变更为nil罢了，更新前后无区别
	if list2 == nil {
		tail.Next = list1
	}
	return head
}

func TestMergeList(t *testing.T) {
	fmt.Println(mergeTwoLists(nil, nil))
}
