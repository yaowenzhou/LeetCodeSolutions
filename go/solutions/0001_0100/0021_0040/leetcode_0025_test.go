// https://leetcode.cn/problems/reverse-nodes-in-k-group/
package solutions

import (
	"fmt"
	"leetcode_solutions_go/solutions/common"
	"testing"
)

func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	curNode := head
	head = &common.ListNode{Next: head}          // 在head前面再增加一个节点，最后返回该节点的next即可
	preNode := head                              // 保存逆转区间的前一个节点
	reverseArr := make([]*common.ListNode, 0, k) // 保存待逆转的节点(当做栈使用)
	for curNode != nil {
		reverseArr = append(reverseArr, curNode) // 待逆转的节点入栈，容量为k
		curNode = curNode.Next
		if len(reverseArr) == k { // 栈慢了，需要进行逆转
			preNode.Next = reverseArr[k-1] // preNode.Next指向数组的最后一个元素
			reverseArr[0].Next = curNode   // reverseArr[0].Next指向curNode(curNode已经被更新过，实际上是当前节点的下一个节点)
			preNode = reverseArr[0]        // 更新preNode为数组的第一个元素，因为在逆转之后，它是下一次逆转的前一个节点
			for i := k - 1; i > 0; i-- {   // 倒序输出(遍历栈)，将节点逆转
				reverseArr[i].Next = reverseArr[i-1]
			}
			// 清空栈
			// 真正的栈是不需要清空，在倒序输出完毕的时候也正好被清空了
			// 此处是一个数组，毕竟不是真正的栈
			reverseArr = reverseArr[0:0]
		}
	}
	return head.Next
}

func TestReverseKGroup(t *testing.T) {
	head := &common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 2,
			Next: &common.ListNode{
				Val: 3,
				Next: &common.ListNode{
					Val: 4,
					Next: &common.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(reverseKGroup(head, 2))
}
