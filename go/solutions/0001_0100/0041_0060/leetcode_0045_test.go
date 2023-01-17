// https://leetcode.cn/problems/jump-game-ii/

package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 解题思路:
// 我自己能想到的就是暴力枚举
// 不过官解给出了更有方案
// 每一步跳跃时，计算落点位置能到达的最远位置
// 这样每次落地都是最优解，到最后自然也是最优解
// 在具体的实现中，我们维护当前能够到达的最大下标位置，记为边界。
// 我们从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
// 在遍历数组时，我们不访问最后一个元素，
// 这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，否则就无法跳到最后一个位置了。如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次「不必要的跳跃次数」，因此我们不必访问最后一个元素。

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestJump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
}
