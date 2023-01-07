// https://leetcode.cn/problems/reverse-nodes-in-k-group/
package solutions

import (
	"fmt"
	"testing"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	curNode := head
	head = &ListNode{Next: head}
	preNode := head // 保存逆转的前一个节点
	reverseArr := make([]*ListNode, 0, k)
	for curNode != nil {
		reverseArr = append(reverseArr, curNode)
		curNode = curNode.Next
		if len(reverseArr) == k {
			preNode.Next = reverseArr[k-1] // preNode.Next指向数组的最后一个元素
			reverseArr[0].Next = curNode   // reverseArr[0].Next指向curNode(curNode已经被更新过，实际上是当前节点的下一个节点)
			preNode = reverseArr[0]        // 更新preNode为数组的第一个元素，因为在逆转之后，它是下一次逆转的前一个节点
			for i := k - 1; i > 0; i-- {
				reverseArr[i].Next = reverseArr[i-1]
			}
			reverseArr = reverseArr[0:0]
		}
	}
	return head.Next
}

func TestReverseKGroup(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(reverseKGroup(head, 2))
}
