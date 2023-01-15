// https://leetcode.cn/problems/number-of-different-subsequences-gcds/
package solutions

import (
	"leetcode_solutions_go/algorithm"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 这是一个很好的解，我自己的算法执行时间太长，就不写了，下面有一个针对它的优化
// https://leetcode.cn/problems/number-of-different-subsequences-gcds/solution/by-lcbin-c0cq/
// 对于数组 nums 的所有子序列，其最大公约数一定不超过数组中的最大值 max。
// 因此我们可以枚举 [1,.. max] 中的每个数x，
// 判断 x 是否为数组 nums 的子序列的最大公约数，如果是，则答案加一。
// 那么问题转换为：判断 x 是否为数组 nums 的子序列的最大公约数
// 我们可以通过枚举 x 的倍数 y，判断 y 是否在数组 nums 中
// 如果 y 在数组 nums 中，则计算 y 的最大公约数 g，如果出现 g=x，则 x 是数组 nums 的子序列的最大公约数。
func countDifferentSubsequenceGCDs(nums []int) (ans int) {
	max := 0
	for _, x := range nums {
		max = algorithm.Max(max, x)
	}
	vis := make([]bool, max+1)
	for _, x := range nums {
		vis[x] = true
	}
	for x := 1; x <= max; x++ {
		g := 0
		for y := x; y <= max; y += x {
			if vis[y] {
				g = algorithm.Gcd(g, y)
				if g == x {
					ans++
					break
				}
			}
		}
	}
	return
}

// 接下来写一个自己的理解
// 首先，官解在枚举过程中使用了数组保存nums
// 但是对于数据量小，但是数字很大的情况下，这种情况无疑会造成很大的内存浪费
// 因此我使用map保存nums,
// map相对于数组的取值效率并不会低多少, 可以做到在时间和空间中取得一个平衡
func countDifferentSubsequenceGCDs1(nums []int) (ans int) {
	numMap := make(map[int]struct{})
	max := nums[0]
	for _, v := range nums {
		numMap[v] = struct{}{}
		if v > max {
			max = v
		}
	}
	ans = len(numMap) // nums最少可以提供len(numMap)个解
	// 开始枚举i,i的最大值为max
	for i := 1; i < max; i++ {
		if _, ok := numMap[i]; ok { // 当i是数组中的某个数时，它必定是一个解，因此直接跳过
			continue
		}
		g := 0
		for j := i; j <= max; j += i { // 枚举i的倍数
			if _, ok := numMap[j]; ok {
				g = algorithm.Gcd(g, j)
				if g == i {
					ans++
					break
				}
			}
		}
	}
	return
}

func TestCountDifferentSubsequenceGCDs(t *testing.T) {
	assert.Equal(t, 7, countDifferentSubsequenceGCDs([]int{5, 15, 40, 5, 6}))
	assert.Equal(t, 7, countDifferentSubsequenceGCDs1([]int{5, 15, 40, 5, 6}))
}
