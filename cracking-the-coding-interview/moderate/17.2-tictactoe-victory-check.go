package main

import (
	"fmt"
)

func main() {
}

func check(board [][]int) bool {
	var first int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if j == 0 && i == 0 {
				first = board[i][j]
			}
			if board[i][j] != first {
				break
			}
		}
	}
}

func checkRows(board [][]int) (int, bool) {
	var first int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if j == 0 && i == 0 {
				first = board[i][j]
			}
			if board[i][j] != first {
				break
			}
			if j == len(board[i])-1 {
				return first, true
			}
		}
	}
	return -1, false
}

func checkCols(board [][]int) (int, bool) {
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board); j++ {
			if j == 0 && i == 0 {
				first = board[j][i]
			}
			if board[j][i] != first {
				break
			}
			if j == len(board)-1 {
				return first, true
			}
		}
	}
	return -1, false
}
