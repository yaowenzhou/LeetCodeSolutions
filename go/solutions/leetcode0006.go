package solutions

func convert(s string, numRows int) string {
	runesIn := []rune(s)
	length := len(runesIn)
	if numRows == 1 || numRows >= length {
		return s
	}
	t := numRows*2 - 2
	runesOut := make([]rune, 0, length)
	// 对于每个周期的每一行，都存在两个字符(第一行和最后一行只有一个字符，但是可以在算法中排除掉，后面代码中有注释)
	// 行枚举
	for i := 0; i < numRows; i++ {
		// 周期枚举，此处使用的是周期的起始索引
		for j := 0; j+i < length; j += t { // j+i表示周期在某一行的第一个字符，必须小于length(索引存在)
			// 周期在每一行的第一个字符可直接追加
			runesOut = append(runesOut, runesIn[i+j])
			if i > 0 && i < numRows-1 && j+t-i < length {
				runesOut = append(runesOut, runesIn[j+t-i])
			}
		}
	}
	return string(runesOut)
}
