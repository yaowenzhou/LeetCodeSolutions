// https://leetcode.cn/problems/merge-k-sorted-lists/
package solutions

import (
	"container/heap"
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

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1] //注意这里，虽然取的是索引n - 1位置的元素，实际上对应的是h[0]，即堆顶元素
	*pq = old[0 : n-1]
	return x
}

// 优先队列实现
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	pq := &PriorityQueue{}
	heap.Init(&PriorityQueue{}) //初始化heap，小根堆(数值越小的越靠前)
	for _, v := range lists {   // 将链表的头结点添加到优先队列
		if v != nil {
			heap.Push(pq, v)
		}
	}
	head := &ListNode{}
	tail := head
	for pq.Len() > 0 { // 只要优先队列中还有链表头结点，就取出节点
		// 由于小根堆保证了优先队列总会弹出当前的最小值
		// 因此我们拿到的永远是剩余链表中头结点最小的那个
		list := heap.Pop(pq).(*ListNode)
		tail.Next = list      // 将取出的链表的头结点追加到待返回的链表中
		tail = tail.Next      // 更新尾结点
		if list.Next != nil { // 如果当前当前链表下一个节点还存在，则将其作为头节点继续放入优先队列
			heap.Push(pq, list.Next)
		}
	}
	return head.Next
}

// 多重循环实现
func mergeKLists2Cycle(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	listsLength := len(lists)
	head := lists[0]
	minNodeIndex := 0
	for i := 0; i < listsLength; i++ {
		if lists[i] == nil {
			continue
		}
		// 此时lists[i不为空]
		if head == nil || lists[i].Val < head.Val { // 此时head也不为空
			head = lists[i]
			minNodeIndex = i
		}
	}
	if head == nil { // 全部都是空节点
		return nil
	}
	lists[minNodeIndex] = lists[minNodeIndex].Next
	if head == lists[0] { // lists[0]为空时，可能发生空指针引用。但是此时head不为空，此分支无法执行
		lists[0] = lists[0].Next
	}
	tail := head // 定义尾结点
	for {
		var minNode *ListNode = nil // 保存当前最小节点
		var minNodeIndex int = -1   // 保存当前最小节点所在链表在切片中的索引
		for i := 0; i < listsLength; i++ {
			if lists[i] == nil {
				continue
			}
			if minNode == nil || lists[i].Val < minNode.Val {
				minNode = lists[i]
				minNodeIndex = i
			}
		}
		// 如果minNode还是空，则说明所有的链表都到底了
		if minNode == nil {
			return head
		}
		// 将最小节点追加到返回链表的尾部
		tail.Next = minNode
		tail = tail.Next
		// 将最小节点从其所在的链表中弹出
		lists[minNodeIndex] = lists[minNodeIndex].Next
	}
}

func TestMergeKLists(t *testing.T) {
	var lists []*ListNode
	lists = append(lists, &ListNode{1, &ListNode{4, &ListNode{5, nil}}})
	lists = append(lists, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	lists = append(lists, &ListNode{2, &ListNode{6, nil}})
	fmt.Println(mergeKLists(lists))
	fmt.Println(mergeKLists2Cycle(lists))
}
