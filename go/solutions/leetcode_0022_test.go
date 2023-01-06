// https://leetcode.cn/problems/generate-parentheses/
package solutions

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) (ret []string) {
	if n == 0 {
		return []string{}
	}
	// 确定左括号的位置
	// 字符串str长度为 2n
	// str[0] 必定为左括号
	// 剩下的 n-1个左括号出现在 [1, 2n-2] 这个闭区间内
	// 排列组合的数量为
	strLen := 2 * n
	var f func(string, int, int)
	// f 对str进行追加字符，i表示第i个字符，leftCnt表示剩余未配对的左括号数量
	f = func(str string, i, leftCnt int) {
		if i == strLen {
			ret = append(ret, str)
			return
		}
		if leftCnt == 0 { // 未配对的左括号数量为0，则只能追加左括号
			f(str+"(", i+1, leftCnt+1)
			return
		} else if leftCnt == strLen-i { // 未配对的左括号数量等于剩余字符串长度，则只能追加右括号
			f(str+")", i+1, leftCnt-1)
			return
		}
		f(str+"(", i+1, leftCnt+1)
		f(str+")", i+1, leftCnt-1)
	}
	f("", 0, 0)
	// fmt.Println(len(ret))
	return ret
}

// 抄的大佬的答案
func generateParenthesis1(n int) (ret []string) {
	if n == 0 {
		return
	}
	var f func(string, int, int)
	f = func(strIn string, i, j int) {
		if i == 0 && j == 0 {
			ret = append(ret, strIn)
			return
		}
		if i > 0 {
			f(strIn+"(", i-1, j)
		}
		if j > i {
			f(strIn+")", i, j-1)
		}
	}
	f("", n, n)
	return
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(len(generateParenthesis(3)))
}

func TestGenerateParenthesis1(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(len(generateParenthesis1(i)))
	}
}
