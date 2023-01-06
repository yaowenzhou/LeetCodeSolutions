// https://leetcode.cn/problems/4sum/
package solutions

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) (ret [][]int) {
	length := len(nums)
	// nums数量小于4，直接返回nil
	if length < 4 {
		return nil
	}
	// 先排序
	sort.Ints(nums)
	for a := 0; a < length-3; a++ { // 外层循环，第一个元素
		if a > 0 && nums[a] == nums[a-1] || nums[a]+nums[length-3]+nums[length-2]+nums[length-1] < target {
			continue
		}
		for b := a + 1; b < length-2; b++ { // 内层循环，第二个元素
			if b > a+1 && nums[b] == nums[b-1] || nums[a]+nums[b]+nums[length-2]+nums[length-1] < target {
				continue
			}
			rest := target - nums[a] - nums[b]
			// 在剩余的元素中，找到nums[c]+nums[d] == target的组合
			for c, d := b+1, length-1; c < d; {
				min := nums[c] + nums[c+1] // 针对第三重循环，nums[c] + nums[c+1]是最小的
				if min > rest {            // 剩下的组合不可能出现四值相加等于target的
					break
				}
				if min == rest {
					ret = append(ret, []int{nums[a], nums[b], nums[c], nums[c+1]})
					break
				}
				max := nums[d] + nums[d-1] // 针对第三重循环，nums[c] + nums[c+1]是最大的
				if max < rest {            // 剩下的组合不可能出现四值相加等于target的
					break
				}
				if max == rest {
					ret = append(ret, []int{nums[a], nums[b], nums[d-1], nums[d]})
					break
				}
				cdSum := nums[c] + nums[d]
				if cdSum > rest {
					for c < d && nums[d] == nums[d-1] {
						d--
					}
					d--
				} else {
					if cdSum == rest {
						ret = append(ret, []int{nums[a], nums[b], nums[c], nums[d]})
					}
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					c++
				}
			}
		}
	}
	return ret
}

func TestFourSum(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
