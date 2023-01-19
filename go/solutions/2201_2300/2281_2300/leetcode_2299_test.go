// https://leetcode.cn/problems/strong-password-checker-ii/
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	specialChars := [128]bool{
		'!': true, '@': true, '#': true, '$': true,
		'%': true, '^': true, '&': true, '*': true,
		'(': true, ')': true, '-': true, '+': true,
	}
	var hasNum, hasLow, hasUp, hasSpecial bool
	if password[0] >= '0' && password[0] <= '9' {
		hasNum = true
	}
	if password[0] >= 'a' && password[0] <= 'z' {
		hasLow = true
	}
	if password[0] >= 'A' && password[0] <= 'Z' {
		hasUp = true
	}
	if specialChars[password[0]] {
		hasSpecial = true
	}
	for i := 1; i < len(password); i++ {
		if password[i-1] == password[i] {
			return false
		}
		if password[i] >= '0' && password[i] <= '9' {
			hasNum = true
		}
		if password[i] >= 'a' && password[i] <= 'z' {
			hasLow = true
		}
		if password[i] >= 'A' && password[i] <= 'Z' {
			hasUp = true
		}
		if specialChars[password[i]] {
			hasSpecial = true
		}
	}
	return hasNum && hasLow && hasUp && hasSpecial
}

func TestStrongPasswordCheckerII(t *testing.T) {
	assert.Equal(t, false, strongPasswordCheckerII("aB3!"))          // 长度不够
	assert.Equal(t, false, strongPasswordCheckerII("ABCDEB3!"))      // 无小写字母
	assert.Equal(t, false, strongPasswordCheckerII("abcdef3!"))      // 无大写字母
	assert.Equal(t, false, strongPasswordCheckerII("ABCdef!^"))      // 无数字
	assert.Equal(t, false, strongPasswordCheckerII("ABCdef12"))      // 无特殊字符
	assert.Equal(t, false, strongPasswordCheckerII("AACdef12!"))     // 连续相同的字符
	assert.Equal(t, true, strongPasswordCheckerII("IloveLe3tcode!")) // 符合要求
}
