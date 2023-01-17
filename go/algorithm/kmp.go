package algorithm

// kmp算法中构造next数组的函数
// j表示next[i-1]
// 如果p[j] = p[i] ==> next[i] = j+1
// 如果p[j] != p[i],则将j迭代为next[j-1]，继续和p[i]进行比较
func getNext(p []rune) []int {
	length := len(p)
	next := make([]int, length)
	j := 0 // p[i-1]的前缀长度
	for i := 1; i < length; i++ {
		for j > 0 && p[j] != p[i] {
			j = next[j-1]
		}
		if p[i] == p[j] {
			j++
			next[i] = j
		}
		// 另一种情况应该将p[j]设为0，但是构造next时，默认值就是0，所以可以不用写
	}
	return next
}

// 查找在s中，从start开始，p第一次出现的索引
func getFirstIndex(s, p []rune, next []int, start int) int {
	sLen, pLen := len(s), len(p)
	if sLen-start < pLen {
		return -1
	}
	for i, j := start, 0; i < sLen; i++ {
		for j > 0 && s[i] != p[j] { // 当j>0且p[i] != p[j]时,将j迭代为next[j-1]
			j = next[j-1]
		}
		if s[i] == p[j] {
			j++
		}
		if j == pLen { // p完全匹配到了
			return i - pLen + 1
		}
	}
	return -1
}

func getAllIndex(s, p []rune, next []int) []int {
	sLen, pLen := len(s), len(p)
	var indexs []int
	for i := 0; i+pLen < sLen; {
		index := getFirstIndex(s, p, next, i)
		if index == -1 {
			break
		}
		i = index + 1
	}
	return indexs
}

// KMP算法(找到p在s中第一次出现的位置)
// @greedy: 是否采取贪心模式，贪心模式将返回p在s中出现的所有起始索引
func Kmp(s, p []rune, greedy bool) (indexs []int) {
	if len(p) > len(s) {
		return
	}
	// 首先构造next
	next := getNext(p)
	if greedy {
		indexs = getAllIndex(s, p, next)
	} else {
		index := getFirstIndex(s, p, next, 0)
		if index != -1 {
			indexs = []int{index}
		}
	}
	return
}

// KmpWithQuestionMark 一个带问号的Kmp算法
// TODO: 目前还不知道怎么实现
// ?表示占一个字符，它可以和任意字符进行匹配
// @greedy: 是否采取贪心模式，贪心模式将返回p在s中出现的所有起始索引
// func KmpWithQuestionMark(s, p []rune, greedy bool) (indexs []int) {
// 	return
// }
