// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
package solutions

import (
	"fmt"
	"testing"
)

func findSubstring(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	// 将s划分为x个单词，有n种划分方法
	// 这是因为，起始位置i必须小于一个字符
	// 如果大于一个字符，则前面可以出现最少一个单词，划分就出了问题
	// 针对划分方法进行遍历
	// 为了保证划分后的单词数量大于等于m,i+m*n <= ls
	// (words中的单词会进行拼接，然后和s进行比较)
	for i := 0; i < n && i+m*n <= ls; i++ {
		// 构建滑动窗口，将当前窗口内的单词频次进行统计
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		// 将words中的单词在differ中减一
		for _, word := range words {
			differ[word]--
		}
		// 滑动窗口,循环开始时，窗口的起点和当前划分方法对的起点相同
		for start := i; start < ls-m*n+1; start += n {
			// 构建窗口时，已经统计过窗口的单词，因此start == i时不再做窗口维护
			if start != i {
				// word为当前移动时进入窗口的单词
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++         // 移入的单词频率+1
				if differ[word] == 0 { // 随时维护单词数量
					delete(differ, word)
				}
				word = s[start-n : start] // 当前移动时移出窗口的单词
				differ[word]--            // 移入的单词频率-1
				if differ[word] == 0 {    // 随时维护单词数量
					delete(differ, word)
				}
			}
			if len(differ) == 0 { // 统计当前窗口的单词数量，如果为0，这说明和words匹配上了
				ans = append(ans, start) // 记录下start
			}
		}
	}
	return
}

func TestSlice(t *testing.T) {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
