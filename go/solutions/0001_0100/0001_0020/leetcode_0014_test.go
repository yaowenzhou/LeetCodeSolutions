// https://leetcode.cn/problems/longest-common-prefix/
package solutions

import (
	"fmt"
	"math"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	var prefix []byte
	// 获取最短字符串长度，最长公共前缀的长度不可能超过长度最短的字符串的长度
	minLength := math.MaxInt
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	for i := 0; i < minLength; i++ {
		var r byte // 保存当前需要遍历的字符
		for _, str := range strs {
			if r == 0 {
				r = str[i]
			}
			// 如果str[i] 和当前遍历的字符不同，则直接返回prefix
			if str[i] != r {
				return string(prefix)
			}
		}
		// 走到这里说明当前遍历的字符符合要求，是公共前缀的一部分
		prefix = append(prefix, strs[0][i]) // 既然是最长公共前缀，那么随便取一个字符串，其第i个字符都和其他字符串的字符一样
	}
	return string(prefix)
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
