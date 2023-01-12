package algorithm

import (
	"golang.org/x/exp/constraints"
)

// Min TODO
func Min[T constraints.Ordered](x T, vs ...T) T {
	min := x
	for _, v := range vs {
		if x > v {
			min = v
		}
	}
	return min
}

// Max TODO
func Max[T constraints.Ordered](x T, vs ...T) T {
	max := x
	for _, v := range vs {
		if x < v {
			max = v
		}
	}
	return max
}

// MaxAbs TODO
func MaxAbs[T constraints.Integer |
	constraints.Float](x T, vs ...T) T {
	max := x
	maxAbs := max
	if maxAbs > 0 { // 为了避免越界，全部转化成负数
		maxAbs = -maxAbs
	}
	for _, v := range vs {
		vAbs := v
		if v > 0 {
			vAbs = -vAbs
		}
		if vAbs < maxAbs {
			max = v
		}
	}
	return max
}

// MinAbs TODO
func MinAbs[T constraints.Integer |
	constraints.Float](x T, vs ...T) T {
	min := x
	minAbs := min
	if minAbs > 0 { // 为了避免越界，全部转化成负数
		minAbs = -minAbs
	}
	for _, v := range vs {
		vAbs := v
		if v > 0 {
			vAbs = -vAbs
		}
		if vAbs > minAbs {
			min = v
		}
	}
	return min
}

// AbsCompare TODO
func AbsCompare[T constraints.Integer |
	constraints.Float](a, b T) int {
	aAbs := a
	bAbs := b
	if a > 0 {
		aAbs = -a
	}
	if b > 0 {
		bAbs = -b
	}
	if aAbs < bAbs {
		return 1
	}
	if aAbs == bAbs {
		return 0
	}
	return -1
}
