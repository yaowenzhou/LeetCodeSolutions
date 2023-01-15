// https://leetcode.cn/problems/min-max-game/

package solutions

import (
	"leetcode_solutions_go/algorithm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minMaxGame(nums []int) int {
	for n := len(nums); n > 1; {
		n >>= 1
		for i := 0; i < n; i++ {
			a, b := nums[i<<1], nums[i<<1|1]
			if i%2 == 0 {
				nums[i] = algorithm.Min(a, b)
			} else {
				nums[i] = algorithm.Max(a, b)
			}
		}
	}
	return nums[0]
}

func TestMinMaxGame(t *testing.T) {
	assert.Equal(t, 1, minMaxGame([]int{1, 3, 5, 2, 4, 8, 2, 2}))
}
