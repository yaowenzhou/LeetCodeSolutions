package solutions

import "math"

func reverse(x int) int {
	a, b := 0, 0
	for x != 0 {
		b = x % 10
		// 判断是否越界
		if a > int(math.MaxInt32)/10 || (a == int(math.MaxInt32)/10 && b > int(math.MaxInt32%10)) {
			return 0
		}
		if a < int(math.MinInt32)/10 || (a == int(math.MinInt32)/10 && b < int(math.MinInt32%10)) {
			return 0
		}
		a = a*10 + b
		x = x / 10
	}
	return a
}
