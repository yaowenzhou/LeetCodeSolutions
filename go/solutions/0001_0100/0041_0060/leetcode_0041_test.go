// https://leetcode.cn/problems/first-missing-positive/
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 解题思路: 最容易想到的是map的方式
// 将数据遍历一遍，全部存储到map
// 然后从1开始枚举，如果某个数没有出现在map中，则该数是缺失的最小正数
// 但是题目要求空间复杂度为O(1)，此方法不合适
// 于是我们可以换个思路，
// 题目没说不可以修改数组那么我们可以利用原来的数组实现
// 对于数组nums,设其长度为N，其缺失的最小正数一定在 [1,N+1]中，如果[1,N]都出现了，则缺失的最小正数为N+1
// 因此我们可以遍历nums, 最初的想法是，如果某个数x在[1-N]中，则，将nums[x-1]的值设为x
// 最后遍历一遍nums，在第一个nums[i] !=nums[i+1]的地方返回，
// 但是这样有个弊端，如果最开始num[2]的位置是5，那么nunms[4(4=5-1)]的数就会被覆盖
// 解决此bug，只需要换个思路，我们使用i来代表一个[1,N]之间的数，至于nums[i]存储的具体值我们并不关心
// 因此，再换个思路，使用正负号在x对应的(x-1)位置进行标记，被打上标记，说明x这个数出现过
// 而nums[x-1]还能拿到原来的值，取绝对值即可
// 这就需要我们给所有<=0的数据全部设置一个不在[1,N]范围的数，例如设置成N+1
// 此时数据全部是正数，然后再遍历一遍数据，对于遍历到的数x，如果满足 1<=|x|<=N，则在x-1处添加负号
// 考虑到数据重重复，因此我们对于已经打上负号的位置可以不用管
// 最后再遍历一遍nums，返回第一个不是负数的索引i+1

func firstMissingPositive1(nums []int) (ret int) {
	n := len(nums)
	// 设正数
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	// 打标记
	for i := 0; i < n; i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if num < n+1 && nums[num-1] > 0 {
			nums[num-1] = -nums[num-1] // 设为负数
		}
	}
	// 查找第一个不为负数的索引
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

// 上面的 "如果某个数x在[1-N]中，则，将nums[x-1]的值设为x" 思路中
// nums[i] = x, 且 x ∈ [1-N]时，将其存放到 x-1 的位置，会导致nums[x-1]中原本的数据消失
// 所以只要我们将nums[x-1]原本的数据保存即可
// 我们将nums[i]=x存放到nums[x-1]处，那么将nums[i]就空出来了，用其来存放nums[x-1]即可
// 所以这是一个交换操作，交换后的数字(nums[x-1])可能处于[1,N]，区间，我们可以继续上述步骤
// 但是这里可能存在一个死循环，如果nums[i] = x = nums[x-1]，此时我们只需要继续下一次循环即可
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func TestFirstMissingPositive(t *testing.T) {
	assert.Equal(t, 2, firstMissingPositive1([]int{3, 4, -1, 1}))
	assert.Equal(t, 2, firstMissingPositive([]int{3, 4, -1, 1}))
}
