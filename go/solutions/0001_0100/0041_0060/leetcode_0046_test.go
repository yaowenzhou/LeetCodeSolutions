// https://leetcode.cn/problems/permutations/
package solutions

import (
	"fmt"
	"testing"
)

func permute(nums []int) (answers [][]int) {
	length := len(nums)
	if length == 0 {
		return
	}
	var insertNum func(int)
	insertNum = func(curIndex int) {
		// 最后一个位置只有一个选择，可以直接append到返回值
		if curIndex == length-1 {
			answers = append(answers, append([]int(nil), nums...))
			return
		}
		for i := curIndex; i < length; i++ {
			// 保持[0,curIndex-1]不变
			// 然后将i和curIndex交换，填充了curIndex
			// curIndex之后的数尚未被选择
			nums[curIndex], nums[i] = nums[i], nums[curIndex]
			insertNum(curIndex + 1)
			nums[curIndex], nums[i] = nums[i], nums[curIndex]
		}
	}
	insertNum(0)
	return
}

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
