// https://leetcode.cn/problems/sudoku-solver/
package solutions

import (
	"fmt"
	"testing"
)

func solveSudoku(board [][]byte) {
	var row [9]int = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var col [9]int = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var are [3][3]int = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	setNum := func(i, j int, num byte) { // 将num添加到row[i]/col[j]/are[i/3][j/3]中
		row[i] |= 1 << num
		col[j] |= 1 << num
		are[i/3][j/3] |= 1 << num
	}
	// 扫描一遍数组,填充已知信息
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			setNum(i, j, board[i][j]-'0')
		}
	}
	// 扫描一遍所有未填的位置，首先填充唯一可能性的位置
	// 即该位置可以确定只能填充某个数
	musk := 0b1111111110
	changed := true
	for changed {
		// 每次遍历空格前，将changed设为false
		// 一旦在遍历过程中，修改了空格的数据，就将changed设为true
		// 直到遍历一遍之后，再也没有唯一可填的数字
		changed = false
		for i := 0; i < 9; i++ {
			for j := 0; j < 0; j++ { // 遍历所有的格子
				if board[i][j] == '.' { // 如果遇到未填数字的格子，就开始试探格子中可填的数字
					for num := byte(1); num <= 9; num++ {
						if (row[i]|1<<num) == musk ||
							(col[j]|1<<num) == musk ||
							(are[i/3][j/3]|1<<num) == musk { // 当num唯一可填时，row[i]||col[j]||board[i][j]的值就和musk相等了
							board[i][j] = '0' + num // 将此数字填充到格子里
							setNum(i, j, num)       // 记录此数字
							changed = true          // 保证最外层循环可以继续执行，继续遍历所有未曾填写数字的格子
							break                   // 唯一可填数字num确定了，最内层试探数字的循环不用执行了
						}
					}
				}
			}
		}
	}
	// 开始枚举无法确定的数字
	var f func(m, n int) bool
	f = func(m, n int) bool {
		for i := m; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					oldRowI, oldColJ, oldAreIJ := row[i], col[j], are[i/3][j/3] // 保存旧值
					for num := byte(1); num <= 9; num++ {                       // 试探数字
						if row[i]&(1<<num) == 0 && // 行不冲突
							col[j]&(1<<num) == 0 && // 列不冲突
							are[i/3][j/3]&(1<<num) == 0 { // 区域不冲突
							// 试着填写该数字，然后进行下一次遍历
							board[i][j] = '0' + num
							setNum(i, j, num) // 记录num
							ok := f(i, j)     // 继续填写下一个格子
							if !ok {          // 回退状态，并进行下一次循环
								row[i], col[j], are[i/3][j/3] = oldRowI, oldColJ, oldAreIJ
								board[i][j] = '.'
							} else { // 此处填写num是ok的，直接返回true
								return ok
							}
						}
					}
					return false
					// 数字试探完了，发现所有的数都冲突，则返回false
				}
			}
		}
		// 所有的空格都填完了就会走到这里，可以直接返回true
		return true
	}
	f(0, 0)
}

func TestSolveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
}
