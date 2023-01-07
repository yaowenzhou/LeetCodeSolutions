package data_structure

import (
	"fmt"
	"testing"
)

type MyInt int

// LessThan for using priority queue
func (i MyInt) LessThan(j PriorityQueueVal) bool {
	return i < j.(MyInt) // 小根堆
	// return i > j.(MyInt) // 大根堆
}

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[MyInt]()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.pq)        // &[1 2]
	fmt.Println(q.IsEmpty()) // false
	fmt.Println(q.Len())     // 2
	fmt.Println(q.Pop())     // 1
	q.Clear()
	fmt.Println(q.Len()) // 0
}
