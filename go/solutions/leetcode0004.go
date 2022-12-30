// https://leetcode.cn/problems/median-of-two-sorted-arrays/
package solutions

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	// totalLeft := (m+n+1)/2 // 左边总比右边多一个或者和右边数量相等
	totalLeft := m + (n-m+1)/2
	median1, median2 := 0, 0
	for left <= right {
		i := left + (right-left)/2   // nums1中划分在左边的元素的数量
		j := totalLeft - i           // nums2中划分在左边的元素的数量
		nums1SpliteL := math.MinInt  // nums1分割线左边的元素(默认取int最小值)
		nums2SplightR := math.MaxInt // nums1分割线右边的元素(默认取int最大值)
		nums2SpliteL := math.MinInt  // nums2分割线左边的元素(默认取int最小值)
		nums2SpliteR := math.MaxInt  // nums2分割线右边的元素(默认取int最大值)
		if i != 0 {                  // 如果nums1有元素划分到左边，则对nums1SpliteL重新进行赋值
			nums1SpliteL = nums1[i-1]
		}
		if i != m { // 如果nums1有元素划分到邮编，则对nums2SplightR重新赋值
			nums2SplightR = nums1[i]
		}
		if j != 0 { // 如果nums2有元素划分到左边，则对nums2SpliteL重新进行赋值
			nums2SpliteL = nums2[j-1]
		}
		if j != n { // 如果nums2有元素划分到左边，则对nums2SpliteL重新进行赋值
			nums2SpliteR = nums2[j]
		}
		if nums1SpliteL <= nums2SpliteR {
			median1 = max(nums1SpliteL, nums2SpliteL)  // 取得左边最大值
			median2 = min(nums2SplightR, nums2SpliteR) // 取得右边最小值
			left = i + 1                               // 分割线右移
		} else {
			right = i - 1 // 分割线左移
		}
	}
	if m%2+n%2 == 1 {
		return float64(median1)
	}
	return float64(median1+median2) / 2.0
}
