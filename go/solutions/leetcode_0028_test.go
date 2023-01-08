// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
package solutions

import (
	"fmt"
	"testing"
)

// 首先搞个函数构造next数组
// 模式串p
// 我们计算p[0:i-1]的最长相同前后缀
// 我们做如下约定:
// 最长相同前后缀中的前缀我们直接用前缀进行称呼
// 最长相同前后缀中的后缀我们直接用后缀进行称呼
// 假定前缀长度为j，则next[i-1]=j
// 并且可以用下面的结构表示字符串
// ^前缀,j,...,后缀,i,...$
// ^表示字符串开始, $表示字符串结束
// 很显然，j是p[0:i-1]的前缀(p[0:j-1])的最后一个字符的索引+1,
// 前缀的首字符索引为0(前缀必定从第一个字符开始), 0=j-j
// i是p[0:i-1]的后缀最后一个字符的索引
// 因此我们可以计算出p[0:i-1]的后缀起始索引为i-j,后缀: p[i-j:i-1]
// 分析1:此时假如 p[i] == p[j],则最长相同前后缀可以+1
// 因为p[0:j-1] == p[i-j:i-1]，p[i]又等于p[j]
// 接下来分析p[i]!=p[j]的情况
// 假设p[i] != p[j]时，存在一个k点，k>j&&p[i]=p[k-1]
// 分析2:首先 k != j+1，因为k=j+1的话，p[i] == p[k-1(==j)]和p[i] != p[j]相矛盾
// 那么有没有可能k>j+1呢,结论是不可能
// 因为存在如下结论：
// p[0:k-1] == p[i-k:i],二者同时去掉最后一个字符，等式依然成立
// p[0:k-2] == p[i-k:i-1] 此长度大于j，这与j的定义违背(j表示p[0:i-1]的最长公共前后缀)
// 因此，k只能小于j,但是它具体在哪里呢？
// 用f(i-1)获取p[0:i-1]的最长公共前后缀，
// 我们在分析1中可知，
// 0---j-1,j,i-j,---i-1,i src
// 					 0  ,---j-1,j,i-j,---i-1,i dest
//							 0, k-1,k,j-k,---j-1,j
// k==p[0:j-1]的前缀长度，而p[0:j-1]的前缀长度已经计算过
// 因此 j 需要迭代为 next[j-1]

func getNext(p []rune) (next []int) {
	length := len(p)
	next = make([]int, length)
	j := 0 // p[i]的前一个字符的前缀长度
	// 第一个字符的前缀为0,make构造时已经初始化为0
	for i := 1; i < length; i++ {
		for j > 0 && p[j] != p[i] {
			j = next[j-1]
		}
		//
		if p[i] == p[j] {
			j++
			next[i] = j
		}
	}
	return
}

func strStr(haystack, needle string) int {
	rlong, rshort := []rune(haystack), []rune(needle)
	longLen, shortLen := len(rlong), len(rshort)
	if shortLen == 0 {
		return 0
	}
	if shortLen > longLen {
		return -1
	}
	next := getNext(rshort)
	for i, j := 0, 0; i < longLen; i++ {
		for j > 0 && rlong[i] != rshort[j] {
			j = next[j-1]
		}
		if rlong[i] == rshort[j] {
			j++
		}
		if j == shortLen {
			return i - shortLen + 1
		}
	}
	return -1
}

func TestStrStr(t *testing.T) {
	fmt.Println(strStr("sadbutsad", "sad"))
}
