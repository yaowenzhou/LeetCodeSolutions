// https://leetcode.cn/problems/longest-valid-parentheses/
package solutions

import (
	"fmt"
	"leetcode_solutions_go/algorithm"
	"testing"
)

func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = algorithm.Max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = algorithm.Max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("(()"))
}
