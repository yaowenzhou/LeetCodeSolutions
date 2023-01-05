// https://leetcode.cn/problems/3sum-closest/
package solutions

import (
	"fmt"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) (ret int) {
	// 为了提高效率，先排个序
	sort.Ints(nums)
	length := len(nums)
	if length <= 3 { // 数组元素数量小于或等于3，直接求和并返回
		for _, v := range nums {
			ret += v
		}
		return ret
	}
	// 首先假定nums[0]+nums[1]+nums[2]就是最接近的数
	ret = nums[0] + nums[1] + nums[2]
	for i := 0; i < length-2; i++ {
		// 固定nums[i]，剩下的数组中找出两数之和与 nums[i] 之和最接近target的结果
		first := i + 1
		last := length - 1
		for first < last {
			twoNum := nums[first] + nums[last] // 提前计算好 nums[first] + nums[last]
			// 优化1:
			// 内循环中 min := nums[i]+nums[first]+nums[first+1]是最小的
			// 如果 min > target, |min - target| 也是内循环中最小的
			// 直接比较 |min-target| < |ret-target| 即可
			min := nums[i] + nums[first] + nums[first+1]
			if min > target {
				if absIntCompare(min-target, ret-target) < 0 {
					ret = min
				}
				break
			}
			// 同理:
			// 内循环中 max := nums[i]+nums[last-1]+nums[last] 是最大的
			// 如果 max < target, |max - target| 也是内循环中最小的
			// 所以如果|max - target| > |ret-target|，则不需要再进行内循环了
			max := nums[i] + nums[last-1] + nums[last]
			if max < target {
				if absIntCompare(max-target, ret-target) < 0 {
					ret = max
				}
				break
			}
			sum := nums[i] + twoNum
			if sum == target {
				return sum // 再没有比它自己更接近它的了
			}
			if absIntCompare(sum-target, target-ret) < 0 {
				ret = sum
			}
			if sum < target {
				for first < last && nums[first] == nums[first+1] {
					first++
				}
				first++
			} else {
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

func TestThreeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
}
