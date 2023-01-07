// priority queue
// 优先队列
package data_structure

import "container/heap"

// import "container/heap"

// PriorityQueueVal the data type using priority queue
// must implement this interface.
// Even for built-in types,
// you should use 'type NewType buildinType'
// to implement this interface
// It determines the priority
// (that is, whether to use a large root heap
// or a small root heap to implement the priority queue)
// 使用优先队列必须实现此接口
// 即使是内置类型你也应该使用`type NewType buildinType`实现此接口
// 它决定了优先级(即使用大根堆还是小根堆实现优先队列)
type PriorityQueueVal interface {
	LessThan(PriorityQueueVal) bool
}

type pq[T PriorityQueueVal] []T

// Len get size of pq
func (pq pq[T]) Len() int {
	return len(pq)
}

// Less TODO
func (pq pq[T]) Less(i, j int) bool {
	return pq[i].LessThan(pq[j])
}

// Swap TODO
func (pq pq[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push TODO
func (pq *pq[T]) Push(x interface{}) {
	*pq = append(*pq, x.(T))
}

// Pop TODO
func (pq *pq[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1] //注意这里，虽然取的是索引n - 1位置的元素，实际上对应的是h[0]，即堆顶元素
	*pq = old[0 : n-1]
	return x
}

type priorityQueue[T PriorityQueueVal] struct {
	pq *pq[T]
}

// NewStack get a priority queue
func NewPriorityQueue[T PriorityQueueVal]() *priorityQueue[T] {
	return &priorityQueue[T]{pq: &pq[T]{}}
}

// Len get the length of priority queue
func (pq *priorityQueue[T]) Len() int {
	return len(*pq.pq)
}

// Push push a value into priority queue
func (pq *priorityQueue[T]) Push(v T) {
	heap.Push(pq.pq, interface{}(v))
}

// Pop pop a value from priority queue
func (pq *priorityQueue[T]) Pop() (v T) {
	return heap.Pop(pq.pq).(T)
}

// Clear clear the priority queue
func (pq *priorityQueue[T]) Clear() {
	*pq.pq = (*pq.pq)[0:0]
}

// IsEmpty tell you if the priority is empty
func (pq *priorityQueue[T]) IsEmpty() bool {
	return len(*pq.pq) == 0
}
