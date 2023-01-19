// https://leetcode.cn/problems/rotate-image/

package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 解题思路
// 假如是一个3*3的矩阵
// [0,0] => [0,2]
// [0,1] => [1,2]
// [0,2] => [2,2]
// 假如是一个4*4的矩阵
// [0,0] => [0,3]
// [0,1] => [1,3]
// [0,2] => [2,3]
// [0,3] => [3,3]
// 可以一圈圈地进行操作，保存当前圈的一条边长的数据，然后进行旋转
// 但是这样做有些麻烦
// 最终参考官解的方案三: 先水平轴翻转，再左上右下作为轴进行对角线翻转
func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func TestRotate(t *testing.T) {
	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	actual := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(actual)
	assert.Equal(t, expected, actual)
}
