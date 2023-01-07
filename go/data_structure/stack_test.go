package data_structure

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	fmt.Println(s)           // &{[1]}
	fmt.Println(s.Pop())     // 1
	fmt.Println(s)           // &{[]}
	fmt.Println(s.Len())     // 0
	fmt.Println(s.IsEmpty()) // true
	s.Push(2)
	fmt.Println(s.IsEmpty()) // false
	s.Clear()
	fmt.Println(s.IsEmpty()) // true
}
