// https://leetcode.cn/problems/combination-sum-ii/
package solutions

import (
	"sort"
	"testing"
)

// 解题思路
// 先将数组排序
// 遍历数组，每个元素可以选中也可以不选中
// 然后加上下一个元素(下一个元素可以选中也可以不选中)
// 当candidates[i]只能不被选中时，可以提前结束递归
// 当被选中的元素集合之和为target时，追加到返回值
// 考虑到解集不能包含重复的组合,此方法还需要去重
// 假如，candidates 经过排序之后
// 某部分的值出现这样的排列 2,2，其索引为i和i+1
// 那么只选中i和只选中i+1结果是一样的，这就是情况重复
// 此时可以参考第39题，我们将其数量统计一下
// 假设数x一共出现了cnt次,那么我们一共可以将x选择[0,cnt]次，左闭右闭
func combinationSum2(candidates []int, target int) (ret [][]int) {
	cntMap := make(map[int]int)
	for i := 0; i < len(candidates); i++ {
		cntMap[candidates[i]]++
	}
	// 使用新的数组进行循环
	candidates = candidates[0:0]
	for k := range cntMap {
		candidates = append(candidates, k)
	}
	sort.Ints(candidates)
	length := len(candidates) // 保存处理之后的数组长度
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
		if curIndex >= length {
			return
		}
		// 性能优化
		// curIndex < length，但是此时target-sum < candidates[curIndex]，说明后续无法再追加数字了，后续的组合达不到要求
		if target-sum < candidates[curIndex] {
			return
		}
		// 切片操作会改变内存，但是递归过程中只会对combination进行追加，因此保存其初始长度即可，后续用于回溯
		// 下面是选中0~cntMap[candidates[curIndex]]的情况
		// 选中过程中sum不断增加，但是sum不能大于target
		for i := 0; sum <= target && i <= cntMap[candidates[curIndex]]; i++ {
			combinationLength := len(combination)
			dfs(curIndex+1, sum, combination)
			sum += candidates[curIndex]
			combination = append(combination[0:combinationLength], candidates[curIndex])
		}
	}
	dfs(0, 0, []int{})
	return
}

func TestCombinationSum2(t *testing.T) {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
}
