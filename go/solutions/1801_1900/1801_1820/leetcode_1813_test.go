// https://leetcode.cn/problems/sentence-similarity-iii/
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 解题思路
// 根据题意，两个句子 sentence1 和 sentence2
// 如果是相似的，那么这两个句子按空格分割得到的字符串数组 words1 和 words2
// 一定能通过往其中一个字符串数组中插入某个字符串数组(可以为空)，得到另一个字符串数组。
// 这个验证可以通过双指针完成。
// i 表示两个字符串数组从左开始，最多有 i 个字符串相同。
// j 表示剩下的字符串数组从右开始，最多有 j 个字符串相同。
// 如果 i+j 正好是某个字符串数组的长度，那么原字符串就是相似的。
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if sentence1 == "" || sentence2 == "" {
		return true
	}
	wordStart := 0
	var words1, words2 []string
	// 构造words1
	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] == ' ' {
			words1 = append(words1, sentence1[wordStart:i])
			wordStart = i + 1
		}
	}
	words1 = append(words1, sentence1[wordStart:])
	// 构造words2
	wordStart = 0
	for i := 0; i < len(sentence2); i++ {
		if sentence2[i] == ' ' {
			words2 = append(words2, sentence2[wordStart:i])
			wordStart = i + 1
		}
	}
	words2 = append(words2, sentence2[wordStart:])
	len1, len2 := len(words1), len(words2)
	if len(words1) > len(words2) { // 为了避免越界，约定words1更小
		words1, words2, len1, len2 = words2, words1, len2, len1
	}
	i, j := 0, 0
	for ; i < len1; i++ {
		if words1[i] != words2[i] {
			break
		}
	}
	if i == len1 {
		return true
	}
	// 从words2的右边开始便利
	for ; j < len1-i; j++ {
		if words1[len1-1-j] != words2[len2-1-j] {
			break
		}
	}
	return i+j == len1
}

func TestAreSentencesSimilar(t *testing.T) {
	assert.Equal(t, true, areSentencesSimilar("My name is Haley", "My Haley"))
	assert.Equal(t, true, areSentencesSimilar("A B C D B B", "A B B"))
}
