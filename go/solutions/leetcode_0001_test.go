// https://leetcode.cn/problems/two-sum/
package solutions

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
