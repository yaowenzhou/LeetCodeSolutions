// https://leetcode.cn/problems/wildcard-matching/
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 解题思路
// 状态转移 => 官解思路1
// https://leetcode.cn/problems/wildcard-matching/solution/tong-pei-fu-pi-pei-by-leetcode-solution/
func isMatch1(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

// 贪心的实现
// TODO: 暂时没有使用KMP的思路
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	// 连续*号等同于一个*，构造一个新的字符串，用于将p中多余的*去掉
	newp := make([]byte, 0, len(p))
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			for ; i < len(p)-1 && p[i+1] == '*'; i++ {
			}
		}
		newp = append(newp, p[i])
	}
	p = string(newp)
	sLen, pLen := len(s), len(p)
	// 1. 将p掐头去尾，遇到星号即停止
	// start表示头的长度，end表示尾的长度
	start, end := 0, 0
	// 如果掐头去尾的过程中，发现匹配不上，即刻返回false
	for ; start < len(p); start++ { // 掐头
		if p[start] == '*' {
			break
		}
		if start == sLen || // p[0:start] 比s长，匹配不上(p[0:start]中不含*)
			(p[start] != '?' && p[start] != s[start]) { // p[start] 和 s[start]匹配不上
			return false
		}
	}
	// 如果s不包含*，即start == pLen，如果s还有剩的，说明不匹配
	if start == pLen && start < sLen {
		return false
	}
	// 头掐完了，现在可以确定s和p存在如下关系
	// [same][start:sLen-1]
	// [same]*[others]
	// end表示尾的长度
	// 所以sLen-start >= end => sLen >= start+end
	for ; pLen-1-end >= start; end++ {
		charp := p[pLen-1-end]
		if charp == '*' {
			break
		}
		if sLen < start+end ||
			(charp != '?' && charp != s[sLen-1-end]) {
			return false
		}
	}
	// 遍历p(p[start]和p[pLen-1-end]是'*')
	// 因此需要遍历p[start+1:pLen-1-end]pLen-1-end是最后一个星号的索引
	// 依次读取单词，*之间的字符串都需要取出，拿到s中进行查找其首次出现的位置
	// 为了避免暴力破解，此处需要写一个kmp算法
	// 为了避免重复计算子串的next(假如某两个星号之间存在相同的字符串)
	// 还需要一个map对next数组进行缓存
	getNext := func(subP string) (next []int) {
		length := len(subP)
		next = make([]int, length)
		j := 0
		for i := 1; i < length; i++ {
			for j > 0 && subP[i] != '?' && subP[j] != '?' && subP[i] != subP[j] {
				j = next[j-1]
			}
			if subP[i] == '?' || subP[j] == '?' || subP[i] == subP[j] {
				j++
				next[i] = j
			}
		}
		return
	}
	// 返回subP在s中startIndex开始，第一次出现的索引
	kmpFind := func(subP string, next []int, startIndex int) int {
		subpLen := len(subP)
		if sLen-startIndex < subpLen {
			return -1
		}
		for i, j := startIndex, 0; i < sLen; i++ {
			for j > 0 && subP[j] != '?' && s[i] != subP[j] {
				j = next[j-1]
			}
			if subP[j] == '?' || s[i] == subP[j] {
				j++
			}
			if j == subpLen {
				return i - subpLen + 1
			}
		}
		return -1
	}
	nextMap := make(map[string][]int)
	// 我们将*之间的子句称之为单词，设wordStart为当前处理的单词在p中的索引
	// sCurIndex表示当前已经匹配上的字符串的长度
	// 从现在开始，end表示p中最后一个*的索引之后的一个位置
	sCurIndex := start
	start, end = start+1, pLen-end
	wordStart := start
	for ; start < end; start++ {
		if p[start] == '*' {
			word := p[wordStart:start]
			next := nextMap[word]
			if next == nil {
				next = getNext(word)
				nextMap[word] = next
			}
			// 查找word第一次出现的索引
			firstIndex := kmpFind(word, next, sCurIndex)
			if firstIndex == -1 {
				return false
			}
			wordStart = start
			sCurIndex = firstIndex + len(word)
			for ; start < end-1 && s[start+1] == '*'; start++ {
			}
		}
	}
	return true
}

func TestIsMatch(t *testing.T) {
	assert.Equal(t, false, isMatch1("aa", "a"))
	// TODO: 贪心算法暂未实现，主要是无法优化带有KMP的字符串查找算法
	assert.Equal(t, true, isMatch("adceb", "*a*b"))
	assert.Equal(t, true, isMatch("aa", "aa"))
	assert.Equal(t, true, isMatch("abcabczzzde", "*abc???de*"))
	assert.Equal(t, false, isMatch("zacabz", "*a?b*"))
}
