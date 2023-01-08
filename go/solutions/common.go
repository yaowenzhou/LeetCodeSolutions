package solutions

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func minAbs(a, b int) int {
	// 先将a，b都转化成正数
	newA := a
	newB := b
	if a < 0 {
		newA = 0 - a
	}
	if b < 0 {
		newB = 0 - b
	}
	if newA <= newB {
		return a
	}
	return b
}

func absIntCompare(a, b int) int {
	newA := a
	newB := b
	if a < 0 {
		newA = 0 - a
	}
	if b < 0 {
		newB = 0 - b
	}
	if newA > newB {
		return 1
	}
	if newA == newB {
		return 0
	}
	return -1
}

func maxAbs(a, b int) int {
	// 先将a，b都转化成正数
	newA := a
	newB := b
	if a < 0 {
		newA = 0 - a
	}
	if b < 0 {
		newB = 0 - b
	}
	if newA >= newB {
		return a
	}
	return b
}
