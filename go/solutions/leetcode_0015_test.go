// https://leetcode.cn/problems/3sum/
package solutions

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	// 为了去重，首先肯定要排序
	sort.Ints(nums)
	length := len(nums)
	if length < 3 { // 凑不齐三元素，直接返回
		return nil
	}
	if nums[0] > 0 || nums[length-1] < 0 { // 都是整数或者都是负数，直接返回
		return nil
	}
	var ret [][]int // 保存结果
	// i != j; i != k; j != k, 同时 nums[i]+nums[j]+nums[k] = 0
	// 假设 i < j < k, 则在完成排序后的数组中nums[i] <= 0
	for i := 0; i < length-2; i++ { // i<j<k，因此i最大取倒数第三个数
		if nums[i] > 0 { // 排序之后的数组，nums[i] > 0,说明剩下的数全是正数，无法达成 nums[i]+nums[j]+nums[k] = 0的条件
			return ret
		}
		// 固定nums[i]，剩下的数组中找出两数之和为 -nums[i]的所有组合
		first := i + 1
		last := length - 1
		rest := 0 - nums[i]
		for first < last {
			if nums[last] < 0 { // 排序之后的数组，nums[last] < 0,说明剩下的数全是正数，无法达成 nums[i]+nums[j]+nums[k] = 0的条件
				break
			}
			sum := nums[first] + nums[last]

			if sum == rest {
				ret = append(ret, []int{nums[i], nums[first], nums[last]})
			}
			if sum <= rest { // sum < rest，first右移，sum = rest时也需要移动，因为需要对比下一个组合（保证可以退出内循环）
				for first < last && nums[first] == nums[first+1] { // 如果出现连续相同的数，此循环执行完之后，first将指向这些相同数的最后一个
					first++
				}
				first++
			} else { // sum > rest, last左移
				for last > first && nums[last] == nums[last-1] { // 如果出现连续相同的数，此循环执行完之后，first将指向这些相同数的第一个
					last--
				}
				last--
			}
		}
		for i < length-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

func TestThreeNum(t *testing.T) {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
