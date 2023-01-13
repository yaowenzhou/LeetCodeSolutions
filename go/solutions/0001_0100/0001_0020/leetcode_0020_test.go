// https://leetcode.cn/problems/valid-parentheses/
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	var charMap [128]byte = [128]byte{'(': ')', '[': ']', '{': '}'}
	stack := make([]byte, length/2) // 括号入栈
	stackLen := len(stack)
	curLen := 0 // 记录左括号数量
	for i := 0; i < length; i++ {
		char := s[i]
		if char == '(' || char == '[' || char == '{' {
			if curLen == stackLen { // 因为要对称，所以在栈满的情况下，不应该出现左括号
				return false
			}
			stack[curLen] = charMap[char]
			curLen++
			continue
		} else {
			if curLen == 0 || char != stack[curLen-1] {
				return false
			} else {
				curLen--
			}
		}
	}
	return curLen <= 0
}
func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()[]{}"))
}
