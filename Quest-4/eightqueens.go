package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, col int) {
	if col == 8 {
		printBoard(board)
		return
	}
	for row := 1; row <= 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board [8]int, col, row int) bool {
	for c := 0; c < col; c++ {
		r := board[c]
		if r == row || r-c == row-col || r+c == row+col {
			return false
		}
	}
	return true
}

func printBoard(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}
