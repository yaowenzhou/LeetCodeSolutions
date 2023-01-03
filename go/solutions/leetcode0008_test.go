// https://leetcode.cn/problems/string-to-integer-atoi/
package solutions

import (
	"fmt"
	"math"
	"testing"
)

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	length := int32(len(s))
	index := int32(0)
	for ; index < length; index++ { // 去除前置空格
		if s[index] != ' ' {
			break
		}
	}
	if index == length { // 如果全部是空格，则需要直接退出,否则后续操作会越界
		return 0
	}
	positive := true
	if s[index] == '-' { // 判断符号
		positive = false
		index++
	} else if s[index] == '+' {
		index++
	}
	num := int32(0)
	for ; index < length; index++ { // 遍历字符串
		if s[index] < '0' || s[index] > '9' { // 非数字字符则退出并根据正负性返回num
			if positive {
				return int(num)
			}
			return int(-1 * num)
		}
		numCur := int32(s[index] - '0') // 获取当前遍历到的数字
		if positive && (num > (math.MaxInt32-numCur)/10 || (num == (math.MaxInt32-numCur)/10 && numCur > math.MaxInt32%10)) {
			return math.MaxInt32
		}
		if !positive && (-1*num < (math.MinInt32+numCur)/10 || (num == (math.MinInt32-numCur)/10 && numCur < math.MinInt32%10)) {
			return math.MinInt32
		}
		num = num*10 + numCur
	}
	if positive {
		return int(num)
	}
	return int(-1 * num)
}

func TestMyAtoi(t *testing.T) {
	fmt.Println("minInt32:", math.MinInt32)
	fmt.Println("maxInt32:", math.MaxInt32)
	fmt.Println(myAtoi(" "))
}
