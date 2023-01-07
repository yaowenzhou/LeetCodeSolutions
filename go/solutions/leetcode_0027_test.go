// https://leetcode.cn/problems/remove-element/
package solutions

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			nums[i] = nums[j]
			j--
		} else {
			i++
		}
	}
	return i
}

func TestRemoveElement(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
