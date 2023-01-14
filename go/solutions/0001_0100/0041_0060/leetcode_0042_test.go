// https://leetcode.cn/problems/trapping-rain-water/
package solutions

import (
	"leetcode_solutions_go/algorithm"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/方法三
func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = algorithm.Max(leftMax, height[left])
		rightMax = algorithm.Max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

func TestTrap(t *testing.T) {
	assert.Equal(t, 4, trap([]int{2, 3, 4, 1, 6, 6, 6, 1, 0, 1}))
}
