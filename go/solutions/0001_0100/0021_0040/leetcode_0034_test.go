// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

package solutions

import (
	"fmt"
	"testing"
)

// 解题思路
// 先比较nums[0]和nums[len(nums)-1]，如果不在范围内，则返回[-1, -1]
// 否则二分递归
// 1: searchRange(nums[0:mid], target)和2: searchRange(nums[mid+1:len(nums-1)])
// 返回结果时，如果1返回[-1,-1]则返回2的结果
// 如果searchRange(nums[0:mid], target)不返回[-1,-1]，则看searchRange(nums[mid+1:len(nums-1)])的结果
// 因为要返回的是nums的索引，所以先定义一个函数用于二分查找，将整个nums作为入参
func binarySearchForSearchRange(nums []int, left, right, target int) []int {
	// 针对nums[left:right]全部相等的一种优化
	if nums[left] == nums[right] {
		if nums[left] == target {
			return []int{left, right}
		}
		return []int{-1, -1}
	}
	// target不在范围内，直接返回 []int{-1, -1}
	if target < nums[left] || target > nums[right] {
		return []int{-1, -1}
	}
	// 由于len(nums) < 10^5，所以不会发生right*2越界int32
	// left < right, left+right就更不可能越界了
	mid := (left + right) / 2
	if nums[mid] < target { // target只可能出现在右区间
		return binarySearchForSearchRange(nums, mid+1, right, target)
	}

	// 此时nums[mid] >= target

	if nums[mid+1] > target { // target只可能出现在左区间
		return binarySearchForSearchRange(nums, left, mid, target)
	}
	// nums[mid+1] <= target
	// nums[mid] <= nums[mid+1]
	// 再加target <= nums[mid]
	// 所以nums[mid+1] <= target <= nums[mid]
	// 有因为nums[mid]<=nums[mid+1]
	// 所以真相只有一个 nums[mid] == target == nums[mid+1]
	// 两边都有target出现过
	ret1 := binarySearchForSearchRange(nums, left, mid, target)
	ret2 := binarySearchForSearchRange(nums, mid+1, right, target)
	return []int{ret1[0], ret2[1]}
}

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 || target < nums[0] || target > nums[length-1] {
		return []int{-1, -1}
	}
	if nums[0] == nums[length-1] {
		return []int{0, length - 1}
	}
	return binarySearchForSearchRange(nums, 0, length-1, target)
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
