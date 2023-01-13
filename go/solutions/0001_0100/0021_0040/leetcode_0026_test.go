// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
package solutions

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	// 因为是有序数组，所以保存curNum即可
	maxLen := len(nums)
	if maxLen < 2 {
		return maxLen
	}
	k := 1
	curNum := nums[0]
	for i := 1; i < maxLen; i++ {
		if nums[i] != curNum {
			curNum = nums[i]
			nums[k] = curNum
			k++
		}
	}
	return k
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
