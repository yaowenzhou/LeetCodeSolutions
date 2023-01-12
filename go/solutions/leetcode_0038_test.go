// https://leetcode.cn/problems/count-and-say/
package solutions

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func countAndSay(n int) string {
	arr := []int{1}
	for i := 2; i <= n; i++ {
		newArr := make([]int, 0)
		curNum := arr[0]
		curCnt := 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == curNum {
				curCnt++
			} else {
				var tmpArr []int
				for curCnt > 0 {
					tmpArr = append(tmpArr, curCnt%10)
					curCnt /= 10
				}
				for m := len(tmpArr) - 1; m >= 0; m-- {
					newArr = append(newArr, tmpArr[m])
				}
				newArr = append(newArr, curNum)
				curCnt = 1
				curNum = arr[j]
			}
		}
		newArr = append(newArr, curCnt, curNum)
		arr = newArr
	}
	var arrIntterfaces []interface{}
	for _, v := range arr {
		arrIntterfaces = append(arrIntterfaces, v)
	}
	str := "%d" + strings.Repeat("%d", len(arr)-1)
	return fmt.Sprintf(str, arrIntterfaces...)
}

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, "1211", countAndSay(4))
}
