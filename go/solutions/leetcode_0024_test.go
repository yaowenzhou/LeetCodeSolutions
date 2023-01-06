// https://leetcode.cn/problems/swap-nodes-in-pairs/
package solutions

import (
	"fmt"
	"testing"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	head = &ListNode{Next: head}
	preNode, curNode := head, head.Next
	for curNode != nil {
		if curNode.Next == nil {
			return head.Next
		}
		// 交换节点i,j需要维护四个值
		// n1(preNode), n2(curNode), n3(curNode.Next), n4(curNode.Next.Next)
		n3 := curNode.Next               // 保存n3
		preNode.Next = n3                // n1的next指向n3
		curNode.Next = curNode.Next.Next // n2的next指向n4
		n3.Next = curNode                // n3的next指向n2
		preNode = curNode                // 刷新preNode
		curNode = curNode.Next           // curNode已经处于n3的位置，n3已经完成了交换，需要将其指向下一个位置
	}
	return head.Next
}

func TestSwapPair(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(swapPairs(head))
}
