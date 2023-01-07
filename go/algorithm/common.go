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
	if maxAbs < 0 {
		maxAbs = 0 - maxAbs
	}
	for _, v := range vs {
		vAbs := v
		if v < 0 {
			vAbs = 0 - vAbs
		}
		if maxAbs < vAbs {
			max = v
		}
	}
	return max
}

// AbsCompare TODO
func AbsCompare[T constraints.Integer |
	constraints.Float](a, b T) int {
	aAbs := a
	bAbs := b
	if a < 0 {
		aAbs = 0 - a
	}
	if b < 0 {
		bAbs = 0 - b
	}
	if aAbs > bAbs {
		return 1
	}
	if aAbs == bAbs {
		return 0
	}
	return -1
}

// MinAbs TODO
func MinAbs[T constraints.Integer |
	constraints.Float](x T, vs ...T) T {
	min := x
	minAbs := min
	if minAbs < 0 {
		minAbs = 0 - minAbs
	}
	for _, v := range vs {
		vAbs := v
		if v < 0 {
			vAbs = 0 - vAbs
		}
		if minAbs > vAbs {
			min = v
		}
	}
	return min
}
