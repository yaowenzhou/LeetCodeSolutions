// https://leetcode.cn/problems/divide-two-integers/
package solutions

import (
	"fmt"
	"math"
	"testing"
)

func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}
	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	isPositive := false
	if dividend > 0 {
		dividend = -dividend
		isPositive = !isPositive
	}
	if divisor > 0 {
		divisor = -divisor
		isPositive = !isPositive
	}
	candidates := []int{divisor}
	// (相当于两根绳子，长的如果减去短的，还是比短的要长，则可以考虑将短的加倍，再进行比较)
	// dividend和y都是负数
	// 长度相减等于 -(dividend)-(-y) = y-dividend
	// 剩余长度y-dividend大于等于y的长度(-y)
	// y-dividend >= -y -> y>=y-dividend
	// y的长度加倍 y = y+y
	// 由于y不断翻倍，且y为负数，所以dividend-y是不断变大的
	// 又因为dividend也是负数，所以在dividend-y溢出之前，y就已经溢出了
	// 所以只要拦截y溢出即可
	// y只有翻倍的时候才会修改，那么当y即将溢出的时候
	// y+y < math.MinInt32 -> y < math.MinInt32 - y
	// dividend > math.MinInt32 -> dividend-y > math.MinInt32 - y > y
	// 根据不等式的传递性可得: dividend-y > y; 这和 dividend-y <= y相矛盾
	// 所以不会出现y溢出的情况
	for y := divisor; y >= dividend-y; {
		y += y
		candidates = append(candidates, y)
	}
	ans := 0
	for i := len(candidates) - 1; i >= 0; i-- {
		// candidates[0] = 1*divisor
		// candidates[1] = 2*divisor
		// candidates[2] = 4*divisor
		// ...
		// candidates[i] = 2^i*divisor = 1<<i*divisor
		// 以 -20/-3举例
		// candidates = []int{-3,-6,-12}
		// candidates[i] >= dividend说明candidates[i]的长度比dividend的长度小，
		// 将candidates[i]所代表的倍数加到ans中，同时dividend减去candidates[i]，用余数判断
		// 由于candidates[i]所代表的倍数为 1<<i*divisor，且i不重复，因此
		// ans 的结果可以写作类似下面的形式
		// 0100+0010+0001,因此ans = ans+1<<i 可以写作 ans |= 1<<i
		if candidates[i] >= dividend {
			ans |= 1 << i
			dividend -= candidates[i]
		}
	}
	if isPositive {
		return -ans
	}
	return ans
}

func TestDivide(t *testing.T) {
	fmt.Println(divide(10, 3))
}
