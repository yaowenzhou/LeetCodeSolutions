// https://leetcode.cn/problems/finding-mk-average/
package solutions

import "github.com/emirpasic/gods/trees/redblacktree"

// 解题思路:
// 经过我的验证
// 使用切片实现队列时，随着不断出入队列
// 切片像一个窗口在一段连续内存上不断滑动
// 在此过程中，已经划过的地方，内存会不断回收
// 因此使用切片来实现队列是可行的

type MKAverage struct {
	lo, mid, hi  *redblacktree.Tree
	q            []int
	m, k, s      int
	size1, size3 int
}

func Constructor(m int, k int) MKAverage {
	lo := redblacktree.NewWithIntComparator()
	mid := redblacktree.NewWithIntComparator()
	hi := redblacktree.NewWithIntComparator()
	return MKAverage{lo, mid, hi, []int{}, m, k, 0, 0, 0}
}

func (mk *MKAverage) AddElement(num int) {
	merge := func(rbt *redblacktree.Tree, key, value int) {
		if v, ok := rbt.Get(key); ok {
			nxt := v.(int) + value
			if nxt == 0 {
				rbt.Remove(key)
			} else {
				rbt.Put(key, nxt)
			}
		} else {
			rbt.Put(key, value)
		}
	}

	if mk.lo.Empty() || num <= mk.lo.Right().Key.(int) {
		merge(mk.lo, num, 1)
		mk.size1++
	} else if mk.hi.Empty() || num >= mk.hi.Left().Key.(int) {
		merge(mk.hi, num, 1)
		mk.size3++
	} else {
		merge(mk.mid, num, 1)
		mk.s += num
	}
	mk.q = append(mk.q, num)
	if len(mk.q) > mk.m {
		x := mk.q[0]
		mk.q = mk.q[1:]
		if _, ok := mk.lo.Get(x); ok {
			merge(mk.lo, x, -1)
			mk.size1--
		} else if _, ok := mk.hi.Get(x); ok {
			merge(mk.hi, x, -1)
			mk.size3--
		} else {
			merge(mk.mid, x, -1)
			mk.s -= x
		}
	}
	for ; mk.size1 > mk.k; mk.size1-- {
		x := mk.lo.Right().Key.(int)
		merge(mk.lo, x, -1)
		merge(mk.mid, x, 1)
		mk.s += x
	}
	for ; mk.size3 > mk.k; mk.size3-- {
		x := mk.hi.Left().Key.(int)
		merge(mk.hi, x, -1)
		merge(mk.mid, x, 1)
		mk.s += x
	}
	for ; mk.size1 < mk.k && !mk.mid.Empty(); mk.size1++ {
		x := mk.mid.Left().Key.(int)
		merge(mk.mid, x, -1)
		mk.s -= x
		merge(mk.lo, x, 1)
	}
	for ; mk.size3 < mk.k && !mk.mid.Empty(); mk.size3++ {
		x := mk.mid.Right().Key.(int)
		merge(mk.mid, x, -1)
		mk.s -= x
		merge(mk.hi, x, 1)
	}
}

func (mk *MKAverage) CalculateMKAverage() int {
	if len(mk.q) < mk.m {
		return -1
	}
	return mk.s / (mk.m - 2*mk.k)
}
