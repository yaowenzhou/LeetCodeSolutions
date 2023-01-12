package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0
	for i < j {
		ans = max(ans, min(height[i], height[j])*(j-i))
		// 矮的索引向另一边靠
		// 因为根据木桶原理，移动高柱的索引会导致底部的长减小，但是高却不可能增加
		// 因此只能是移动矮柱的索引
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
