package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 解题思路
// 首先题目给出一个下标从 0 开始长度为 n 的非负整数数组 nums，
// 并给出「好下标对」的定义——对于某一个下标对 (i,j)，
// 0 <= i < j < n，若满足：
// nums[i] + rev(nums[j]) = nums[j] + rev(nums[i])		(1)
// 则该下标对为「好下标对」。
// 现在我们设：f(i) = nums[i] - rev(nums[i])
// 则表达式 (1) 可以等价于:
// f(i) = f(j)											(2)
// 那么我们从左到右遍历数组nums
// 并在遍历的过程用「哈希表」统计每一个位置 i，
// 0 <= i < n 的 f(i) 的个数，
// 则对于位置 j，0 <= j < n
// 以 j 结尾的「好下标对」的个数为此时「哈希表」中 f(j) 的数目。
// 那么我们只需要在遍历时同时统计以每一个位置为结尾的「好下标对」数目即可。

func countNicePairs(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		rev := 0
		for x := num; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		ans += cnt[num-rev]
		cnt[num-rev]++
	}
	return ans % (1e9 + 7)
}

func TestCountNicePairs(t *testing.T) {
	assert.Equal(t, 4, countNicePairs([]int{13, 10, 35, 24, 76}))
}
