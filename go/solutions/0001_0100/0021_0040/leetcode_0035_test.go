// https://leetcode.cn/problems/search-insert-position/
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 定义一个二分查找函数
func binarySearchForSearchInsert(nums []int, left, right, target int) int {
	if target < nums[left] {
		return left
	}
	if target > nums[right] {
		return right + 1
	}
	mid := (left + right) / 2
	if target == nums[mid] {
		return mid
	}
	if target < nums[mid] {
		return binarySearchForSearchInsert(nums, left, mid, target)
	}
	return binarySearchForSearchInsert(nums, mid+1, right, target)
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 || nums[0] >= target {
		return 0
	}
	if nums[length-1] == target {
		return length - 1
	}
	if nums[length-1] < target {
		return length
	}
	return binarySearchForSearchInsert(nums, 0, length-1, target)
}

func TestSearchInsert(t *testing.T) {
	ret := searchInsert([]int{1, 3, 5, 6}, 5)
	assert.Equal(t, 2, ret, "searchInsert error")
}
