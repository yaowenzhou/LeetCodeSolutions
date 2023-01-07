package algorithm

func getStepLength(runesExpend []rune, left, right int) int {
	for ; left >= 0 && right < len(runesExpend) && runesExpend[left] == runesExpend[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) >> 1
}

// 马拉车算法
func Manacher(s string) string {
	// 1. 构造扩展字符串
	runes := []rune(s)
	length := len(runes)
	runesExpend := make([]rune, length*2+1)
	for i := 0; i < length; i++ {
		runesExpend[2*i+1] = runes[i]
	}
	// 记录最大回文子串的起始位置和终点位置
	start, end := 0, -1
	// 记录以每一个字符作为中心能扩展的臂长
	armLenArr := make([]int, len(runesExpend))
	// 记录当前能回文的最右端对应的center
	// 记录当前能回文的最右端索引
	center, maxRight := -1, -1
	for i := 0; i < len(runesExpend); i++ {
		var curArmLen int
		if i > maxRight { // 中心扩散
			curArmLen = getStepLength(runesExpend, i, i)
		} else {
			// i <= maxRight， 则根据center的回文特性，
			// minLeft .... mirrorI .... center .... i .... maxRight
			// 可以保证 i-min(mirrorI的臂长, maxRight-i) .... i .... i+min(mirrorI的臂长, maxRight-i)这一段必定是回文的
			mirrorI := center*2 - i
			minArmLen := Min(armLenArr[mirrorI], maxRight-i)
			curArmLen = getStepLength(runesExpend, i-minArmLen, i+minArmLen)
		}
		armLenArr[i] = curArmLen    // 记录下当前位置的臂长
		if i+curArmLen > maxRight { // 更新center和maxRight
			center = i
			maxRight = i + curArmLen
		}
		if curArmLen*2+1 > end-start {
			start, end = i-curArmLen, i+curArmLen
		}
	}
	subRunes := make([]rune, (end-start+1)>>1)
	cnt := 0
	for i := start; i < end+1; i++ {
		if runesExpend[i] != 0 {
			subRunes[cnt] = runesExpend[i]
			cnt++
		}
	}
	return string(subRunes[0:cnt])
}
