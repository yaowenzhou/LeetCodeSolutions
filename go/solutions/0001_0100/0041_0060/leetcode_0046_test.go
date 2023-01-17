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
	var insertNum func([]int, int)
	insertNum = func(answer []int, curIndex int) {
		// 最后一个位置只有一个选择，可以直接append到返回值
		if curIndex == length-1 {
			answers = append(answers, answer)
		}
		for i := curIndex; i < length; i++ {
			answerNew := make([]int, 0, length)
			for j := 0; j < length; j++ {
				answerNew = append(answerNew, answer[j])
			}
			// 保持[0,curIndex-1]不变
			// 然后将i和curIndex交换，填充了curIndex
			// curIndex之后的数尚未被选择
			answerNew[curIndex], answerNew[i] = answerNew[i], answerNew[curIndex]
			insertNum(answerNew, curIndex+1)
		}
	}
	insertNum(nums, 0)
	return
}

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
