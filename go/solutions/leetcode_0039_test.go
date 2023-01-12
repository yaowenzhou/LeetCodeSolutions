// https://leetcode.cn/problems/combination-sum/

package solutions

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 解题思路：
// 首先遍历candidates
// 针对每一个元素candidates[i]，首先设置自己的倍数分别为0,1,2,3,4...
// 然后分别加上下一个元素的0,1,2,3...倍
// 再继续加上下一个元素的0,1,2,3...倍
// 一直到最后
// 如果等于target，就将这些数填充到ret
// 最后返回整个ret
// 由于元素是从小到大排列，因此不会出现重复

func combinationSum(candidates []int, target int) (ret [][]int) {
	// 题目没说candidates是有序的，因此先排个序
	sort.Ints(candidates)
	length := len(candidates) // 保存长度，后面用得到，提高效率
	// 定义递归函数
	var dfs func(int, int, []int)
	// curIndex是当前要操作的元素的索引，combination是当前的数字组合，sum是当前数字组合之和(主要是为了减少计算，提高效率)
	dfs = func(curIndex, sum int, combination []int) {
		if sum == target { // 当前的组合刚好等于target，则记录此组并返回
			var ints []int
			ints = append(ints, combination...)
			ret = append(ret, ints)
			return
		}
		if curIndex == length {
			return
		}
		// 性能优化
		// curIndex < length，但是此时target-sum < candidates[curIndex]，说明后续无法再追加数字了，后续的组合达不到要求
		if target-sum < candidates[curIndex] {
			return
		}
		// 切片操作会改变内存，但是递归过程中只会对candidates进行追加，因此保存其初始长度即可，后续用于回溯
		for ; sum <= target; sum += candidates[curIndex] {
			combinationLength := len(combination)
			// 从0开始，所以直接调用一遍
			dfs(curIndex+1, sum, combination)
			combination = combination[0:combinationLength]
			combination = append(combination, candidates[curIndex])
		}
	}
	dfs(0, 0, []int{})
	return
}

func TestCombinationSum(t *testing.T) {
	ret := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.Equal(t, []int{7}, ret[0])
	assert.Equal(t, []int{2, 2, 3}, ret[1])
}
