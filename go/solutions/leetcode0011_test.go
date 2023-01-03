package solutions

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0
	for i < j {
		ans = max(ans, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
