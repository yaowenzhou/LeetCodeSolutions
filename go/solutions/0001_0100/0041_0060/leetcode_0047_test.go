// https://leetcode.cn/problems/permutations-ii/

package solutions

import (
	"fmt"
	"testing"
)

func permuteUnique(nums []int) (answers [][]int) {
	length := len(nums)
	if length == 0 {
		return
	}
	numsMap := make(map[int]int)
	for i := 0; i < length; i++ {
		numsMap[nums[i]]++
	}
	nums = nums[0:0] // nums用来存储插入到返回值里面的临时容器
	var insertNum func(int)
	insertNum = func(curIndex int) { // curIndex表征插入了多少数字
		// 数字填完了, 追加到返回值
		if curIndex == length {
			answers = append(answers, append([]int(nil), nums...))
			return
		}
		for num, cnt := range numsMap {
			if cnt == 0 {
				continue
			}
			numsMap[num]--
			nums = append(nums, num)
			insertNum(curIndex + 1)
			nums = nums[:curIndex]
			numsMap[num]++ // 用完数字之后再放回去
		}
	}
	insertNum(0)
	return
}

func TestPermuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{2, 2, 1, 1}))
}
