// https://leetcode.cn/problems/finding-the-users-active-minutes/
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	mp := map[int]map[int]bool{}
	for _, p := range logs {
		id, t := p[0], p[1]
		if mp[id] == nil {
			mp[id] = map[int]bool{}
		}
		mp[id][t] = true
	}
	ans := make([]int, k+1)
	for _, m := range mp {
		ans[len(m)]++
	}
	return ans[1:]
}

func TestFindingUsersActiveMinutes(t *testing.T) {
	assert.Equal(t, []int{0, 2, 0, 0, 0},
		findingUsersActiveMinutes(
			[][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}},
			5))
}
