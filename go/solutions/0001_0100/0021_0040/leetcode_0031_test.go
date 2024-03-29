// https://leetcode.cn/problems/next-permutation/
package solutions

import (
	"fmt"
	"testing"
)

// 解题思路
// “下一个排列” 的定义是：给定数字序列的字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
// 我们可以将该问题形式化地描述为：给定若干个数字，将其组合为一个整数。
// 如何将这些数字重新排列，以得到下一个更大的整数。如 123 下一个更大的数为 132。
// 如果没有更大的整数，则输出最小的整数。
// 以 1,2,3,4,5,6 为例，其排列依次为：
// 123456
// 123465
// 123546
// ...
// 654321
// 可以看到有这样的关系：123456 < 123465 < 123546 < ... < 654321。
// 算法推导
// 我们希望下一个数 比当前数大，这样才满足 “下一个排列” 的定义。
// 1. 因此只需要 将后面的「大数」与前面的「小数」交换，就能得到一个更大的数。
// 此要求需要保证交换时，后面的数一定要大于前面的数
// 比如 123456，将 5 和 6 交换就能得到一个更大的数 123465。
// 我们还希望下一个数 增加的幅度尽可能的小
// 这样才满足“下一个排列与当前排列紧邻“的要求。为了满足这个要求，我们需要：
// 在 尽可能靠右的低位 进行交换，需要 从后向前 查找
// 将一个 尽可能小的「大数」 与前面的「小数」交换。比如 123465，下一个排列应该把 5 和 4 交换而不是把 6 和 4 交换
// 将「大数」换到前面后，需要将「大数」后面的所有数 重置为升序，升序排列就是最小的排列。
// 以 123465 为例：首先按照上一步，交换 5 和 4，得到 123564；然后需要将 5 之后的数重置为升序，得到 123546。
// 显然 123546 比 123564 更小，123546 就是 123465 的下一个排列
// 以上就是求 “下一个排列” 的分析过程。
// 下面是算法描述
// 1. 从后向前 查找第一个 相邻升序 的元素对 (i,j)，满足 A[i] < A[j]。此时 [j,end) 必然是降序
// 2. 在 [j,end) 从后向前 查找第一个满足 A[i] < A[k] 的 k。A[i]、A[k] 分别就是上文所说的「小数」、「大数」
// 3. 将 A[i] 与 A[k] 交换
// 4. 可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
// 如果在步骤 1 找不到符合的相邻元素对，说明当前 [begin,end) 为一个降序顺序，则直接跳到步骤 4

func nextPermutation(nums []int) {
	maxIndex := len(nums) - 1
	if maxIndex <= 0 {
		return
	}
	i, j, k := maxIndex-1, maxIndex, maxIndex
	// find: A[i]<A[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 { // 不是最后一个排列
		// find: A[i]<A[k]
		for nums[i] >= nums[k] {
			k--
		}
		// swap A[i], A[k]
		nums[i], nums[k] = nums[k], nums[i]
	}
	// reverse A[j:end]
	for i, j := j, maxIndex; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
