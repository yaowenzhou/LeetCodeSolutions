// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
package solutions

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) (combinations []string) {
	var num2Chars [][]byte = [][]byte{
		{'a', 'b', 'c'},      // 数字2
		{'d', 'e', 'f'},      // 数字3
		{'g', 'h', 'i'},      // 数字4
		{'j', 'k', 'l'},      // 数字5
		{'m', 'n', 'o'},      // 数字6
		{'p', 'q', 'r', 's'}, // 数字7
		{'t', 'u', 'v'},      // 数字8
		{'w', 'x', 'y', 'z'}, // 数字9
	}
	if digits == "" { // 空字符串直接返回
		return []string{}
	}
	charCnts := []int{3, 3, 3, 3, 3, 4, 3, 4} // 每个数字对应的字符数量(因为需要多次使用，因此提前保存)
	retLen := 1
	for _, v := range digits { // 计算好返回值的数组长度，避免append时无谓的复制
		retLen *= charCnts[v-'2']
	}
	combinations = make([]string, 0, retLen)
	digitsLen := len(digits) // 后续需要多次使用，因此提前计算保存
	var f func(int, string)
	f = func(index int, strIn string) {
		if index == digitsLen {
			combinations = append(combinations, strIn)
		} else {
			digit := digits[index]        // 取数字
			chars := num2Chars[digit-'2'] // 数字的字符列表
			// charCnt := charCnt[digit-'2'] -> 取数字的字符数量
			for i := 0; i < charCnts[digit-'2']; i++ {
				f(index+1, strIn+string(chars[i])) // 这里使用+拼接字符串可以完美实现字符串的复制，利用了go中字符串不可变的特性
			}
		}
	}
	f(0, "")
	return combinations
}

func letterCombinations1(digits string) (combinations []string) {
	var num2Chars [][]byte = [][]byte{
		{'a', 'b', 'c'},      // 数字2
		{'d', 'e', 'f'},      // 数字3
		{'g', 'h', 'i'},      // 数字4
		{'j', 'k', 'l'},      // 数字5
		{'m', 'n', 'o'},      // 数字6
		{'p', 'q', 'r', 's'}, // 数字7
		{'t', 'u', 'v'},      // 数字8
		{'w', 'x', 'y', 'z'}, // 数字9
	}
	if digits == "" { // 空字符串直接返回
		return []string{}
	}
	digitsLen := len(digits)                  // 后续需要多次使用，因此提前计算保存
	charCnts := []int{3, 3, 3, 3, 3, 4, 3, 4} // 每个数字对应的字符数量(因为需要多次使用，因此提前保存)
	// 用一个数组(combinationsTemp)盛装整个计算过程中的所有字符串
	// tmpLen表示这个字符串的最大长度，用于make，以避免append的扩容复制
	// retLen用于计算最终返回的字符串的长度，同时也可以计算tmpLen
	retLen, tmpLen := 1, 1
	for _, v := range digits { // 计算好返回值的数组长度，避免append时无谓的复制
		// 按照digits中数字的顺序，对当前字符串挨个增加该数字代表的字符，并添加到临时数组中
		// 初始字符串设为""
		// 例如 digits = "27"
		// 则最开始 combinationsTemp = []string{""}
		// 第一次循环时，2对应的字符有三个'a', 'b', 'c'
		// 则取出combinationsTemp中的所有字符串，每个字符串挨个追加'a', 'b', 'c'并添加到combinationsTemp中
		// 第一次循环完毕时，combinationsTemp = []string{"", "a", "b", "c"}
		// 此时假如只有一个数字2，则我们应该返回{"a", "b", "c"}
		// 但是还存在第2个数字7,7对应的字符有4个{'p', 'q', 'r', 's'}
		// 那么第二次循环时，我们应该针对第一次循环中添加的这些字符串进行遍历，
		// 每个字符串分别追加'p', 'q', 'r', 's'，并将结果追加到combinationsTemp
		// 如果只有digits = "27"，那么我们就可以返回第二次循环中追加的这些字符串了
		// 根据以上规律，我们很容易得出tmpLen的规律
		// 它等于temLen+每次外层循环(digits中的数字遍历)时新增加的字符串
		retLen *= charCnts[v-'2']
		tmpLen += retLen
	}
	combinationsTemp := make([]string, 1, tmpLen)
	start := 0                       // 标记当前的起始字符串
	for i := 0; i < digitsLen; i++ { // 循环digits的每一个字符
		curNum := digits[i] - '2'
		lenTmp := len(combinationsTemp) - start // 记录字符串数量
		for j := 0; j < lenTmp; j++ {           // 从当前头结点head循环遍历到尾部节点
			for _, char := range num2Chars[curNum] { // 第三重循环: 以当前数字所对应的字符列表进行循环
				combinationsTemp = append(combinationsTemp, combinationsTemp[start]+string(char))
			}
			// 第三重循环之后，start后移一位
			start++
		}
	}
	return combinationsTemp[start:]
}

// Running tool: E:\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkLetterCombinations$ solutions/solutions -count=1
// goos: windows
// goarch: amd64
// pkg: solutions/solutions
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkLetterCombinations
// BenchmarkLetterCombinations-12
//
//	1        230457231800 ns/op      221741125336 B/op       2790750982 allocs/op
//					 132226645700 // 使用make提前分配好返回值的长度后提高的效率
//					 122860589900
//
// PASS
// ok      solutions/solutions     230.778s
// > Test run finished at 1/5/2023, 10:43:19 PM <
func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func TestLetterCombinations1(t *testing.T) {
	fmt.Println(letterCombinations1("23"))
}

func BenchmarkLetterCombinations(b *testing.B) {
	for i := 0; i < 10000; i++ {
		letterCombinations("7897457845")
	}
}

func BenchmarkLetterCombinations1(b *testing.B) {
	for i := 0; i < 10000; i++ {
		letterCombinations1("7897457845")
	}
}
