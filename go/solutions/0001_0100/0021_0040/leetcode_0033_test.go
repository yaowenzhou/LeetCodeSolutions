// https://leetcode.cn/problems/search-in-rotated-sorted-array/
package solutions

import (
	"fmt"
	"testing"
)

// 解题思路
// 1. 设数组长度为length
// 旋转后的数组是如下结构
// -----k,k+1----(0<=k<=length-1)
// [0,k]升序, [k+1,length-1]也是升序
// [0,k] 中的每一个数都大于 [k+1,length-1]
// 则可以使用二分法将数组分为两部分
// [0,mid-1] [mid, length-1]; 其中mid等于length/2
// 通过比较nums[mid]和nums[length-1]的大小可以将数组分成有序和无序两个部分
// 判定target可能在在有序部分中，则直接对有序部分二分查找
// 判定target可能在无序部分，则继续对无序部分进行二分，并重复这个步骤

// 先定义一个二分查找函数
func binarySearch(nums []int, left, right, target int) int {
	for mid := (left + right) / 2; ; mid = (left + right) / 2 {
		// left小于等于right的情况下，mid+1<=right
		// 先判断target是否等于mid，如果等于，则直接返回mid
		// 如果不等于，此时有可能开始循环的时候left==right
		// 或者上次循环时left=right-1，而到了这次循环left正好等于right
		// 这时候，mid == left == right
		// 由于target==nums[mid]的情况已经判定过
		// 因此如果检测到left==right可以直接返回-1
		if target == nums[mid] { // 正好等于nums[mid]则直接返回
			return mid
		}
		if left == right {
			return -1
		}
		if target > nums[mid] { // target可能位于右区间，更新left
			left = mid + 1
		} else { // target可能位于左区间，更新right
			right = mid
		}
	}
}

// 无序查找函数
func disOrderSearch(nums []int, left, right, target int) int {
	if left == right {
		if target == nums[left] {
			return left
		} else {
			return -1
		}
	}
	mid := (left + right) / 2
	// nums[left, right]可以分为两个部分 nums[left,mid]和nums[mid+1,right]
	// 此函数的调用中left<=right是默认条件，因此mid>=left
	// left==right在前面已经排除了，所以mid<right，mid+1<=right
	if nums[left] <= nums[mid] { // nums[left,mid]是有序的
		if target >= nums[left] && target <= nums[mid] { // target位于有序区间，对有序区间使用二分查找
			return binarySearch(nums, left, mid, target)
		}
		// target不位于有序区间，有可能位于另一个区间，对另一个区间递归调用无序查找函数
		return disOrderSearch(nums, mid+1, right, target)
	}
	// nums[left,mid]是无序的,那么nums[mid+1,right]必定有序
	if target >= nums[mid+1] && target <= nums[right] { // target处于有序区间
		return binarySearch(nums, mid+1, right, target)
	}
	return disOrderSearch(nums, left, mid, target) // nums不在有序区间
}

func search(nums []int, target int) int {
	return disOrderSearch(nums, 0, len(nums)-1, target)
}

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
