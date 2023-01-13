// https://leetcode.cn/problems/rearrange-characters-to-make-target-string/
package solutions

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rearrangeCharacters(s string, target string) int {
	// 保存target中各个字符的数量
	// 以字符的值作为下标，保存的是该字符出现的数量
	// 题目中的字符都是ascii中的可见字符，因此都处于[0,127]区间
	var tCharCnt [128]int
	// 统计target中各个字符的数量
	for i := 0; i < len(target); i++ {
		tCharCnt[target[i]]++
	}
	// 保存s中各个字符的数量
	var sCharCnt [128]int
	for i := 0; i < len(s); i++ {
		sCharCnt[s[i]]++
	}
	// 求取最小倍数
	ret := math.MaxInt
	for k, v := range tCharCnt {
		if v != 0 {
			times := sCharCnt[k] / v
			if times < ret {
				ret = times
			}
		}
	}
	return ret
}

func TestRearrangeCharacters(t *testing.T) {
	assert.Equal(t, 2, rearrangeCharacters("ilovecodingonleetcode", "code"))
}
