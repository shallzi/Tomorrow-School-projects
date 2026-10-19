package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solveQueens(0, &board)
}

func solveQueens(col int, board *[8]int) {
	if col == 8 {
		printBoard(board)
		return
	}
	for row := 1; row <= 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solveQueens(col+1, board)
		}
	}
}

func isSafe(board *[8]int, col, row int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || abs(board[i]-row) == abs(i-col) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}
