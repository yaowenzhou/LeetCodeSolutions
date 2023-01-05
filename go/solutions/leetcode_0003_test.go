// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package solutions

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	rs := []rune(s)
	maxLength, start, end := 0, 0, 0
	var runeSlice = [256]int{}   // 保存字符的索引+1
	for ; end < len(rs); end++ { // 终止字符不断右移
		if runeSlice[rs[end]] >= start+1 { // 终止字符索引+1大于等于start+1，说明出现过
			if end-start+1 > maxLength { // 更新字符串最大长度
				maxLength = end - start
			}
			start = runeSlice[rs[end]] // 更新start为终止字符上次出现的索引+1
		}
		runeSlice[int(rs[end])] = end + 1 // 记录当前字符的索引
	}
	if end-start+1 > maxLength {
		maxLength = end - start
	}
	return maxLength
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
