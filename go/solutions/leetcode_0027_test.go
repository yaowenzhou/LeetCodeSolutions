// https://leetcode.cn/problems/remove-element/
package solutions

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	// 双指针, i是当前需要处理的值，j是最后一个未处理的值
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val { // 如果nums[i]是需要移除的值
			nums[i] = nums[j] // 则将nums[j]赋值给nums[i]
			j--               // 同时j左移(nums[j]的值已经给了nums[i])
		} else {
			i++ // 如果nums[i]不是需要移除的值，则i右移
		}
	}
	// 上面的循环中，j始终指向未处理的值，直到i==j时，将j给处理掉
	// i始终指向已经处理过的值得下一个索引，因此最后返回i即可
	return i
}

func TestRemoveElement(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
