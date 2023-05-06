package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Board [9][9]int

func format_print(b Board) string {
	var buff bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buff.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buff.WriteString("|")
			}
			buff.WriteString(strconv.Itoa(b[i][j]))
		}
		buff.WriteString("|\n")
	}
	buff.WriteString("+---+---+---+\n")
	return buff.String()
}

// 0 → 未入力
// 1~9 → 入力済み
// countは0~9が入力された回数を表す
func isDuplicated(count [10]int) bool {
	for k, v := range count {
		// 0は未入力状態なので常に許容する
		if k == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(board Board) bool {
	for i := 0; i < 9; i++ {
		var count [10]int
		for j := 0; j < 9; j++ {
			count[board[i][j]]++
		}
		if isDuplicated(count) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		var count [10]int
		for j := 0; j < 9; j++ {
			count[board[j][i]]++
		}
		if isDuplicated(count) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var count [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					count[board[row][col]]++
				}
				if isDuplicated(count) {
					return false
				}
			}
		}
	}
	return true
}

func isSolved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func backtrack(b *Board) bool {
	if isSolved(*b) {
		fmt.Printf("%v", *b)
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					b[i][j] = candidate

					if verify(*b) {
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func main() {
	board := Board{
		{0, 0, 0, 9, 0, 0, 7, 4, 0},
		{8, 0, 1, 0, 0, 0, 0, 0, 0},
		{6, 0, 7, 0, 0, 0, 5, 0, 3},
		{0, 3, 0, 0, 4, 0, 0, 5, 0},
		{5, 0, 0, 0, 8, 0, 3, 2, 0},
		{4, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 5, 0, 0, 0, 6, 0, 0, 0},
		{0, 0, 0, 0, 3, 0, 0, 1, 9},
		{1, 7, 8, 0, 0, 0, 0, 0, 0},
	}
	backtrack(&board)
}
