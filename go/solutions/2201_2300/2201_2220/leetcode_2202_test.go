// https://leetcode.cn/problems/maximize-the-topmost-element-after-k-moves/
package solutions

import (
	"leetcode_solutions_go/algorithm"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 如果 len(nums) = 1，k 是偶数，那么栈里存在 1个元素；如果 k 是奇数，那么栈里一定没有任何元素，返回 -1。
// 否则，栈顶元素可以是：
// nums 的前 k-1个数的最大值；
// nums 的第 k+1个数（如果有，此时出栈前 k 个数即可）；
// 注意，nums 的第 k 个数永远不会出现在栈顶。

func maximumTop(nums []int, k int) int {
	length := len(nums)
	if length == 0 { // nums为空直接返回-1
		return -1
	}
	if length == 1 && k&1 == 1 {
		return -1
	}
	res := math.MinInt
	for i := 0; i < length && i < k-1; i++ { // 前k-1个数或者所有数的最大值，取决于k和length哪个更小
		res = algorithm.Max(res, nums[i])
	}
	if k < length {
		res = algorithm.Max(res, nums[k]) // 和第k+1个数进行比较
	}
	return res
}

func TestMaximumTop(t *testing.T) {
	assert.Equal(t, 94, maximumTop([]int{
		35, 43, 23, 86, 23, 45, 84, 2, 18, 83,
		79, 28, 54, 81, 12, 94, 14, 0, 0, 29,
		94, 12, 13, 1, 48, 85, 22, 95, 24, 5,
		73, 10, 96, 97, 72, 41, 52, 1, 91, 3,
		20, 22, 41, 98, 70, 20, 52, 48, 91, 84,
		16, 30, 27, 35, 69, 33, 67, 18, 4, 53,
		86, 78, 26, 83, 13, 96, 29, 15, 34, 80,
		16, 49,
	}, 15))
	assert.Equal(t, -1, maximumTop([]int{2}, 1))
}
