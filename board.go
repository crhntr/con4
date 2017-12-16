package main

import (
	"errors"
	"fmt"
)

var (
	errCollumnFull = errors.New("collumn full")
)

func newBoard(cols, rows int) [][]byte {
	board := make([][]byte, cols)
	for col := range board {
		board[col] = make([]byte, rows)
	}
	return board
}

func drop(board [][]byte, val byte, col int) error {
	if col < 0 || col >= len(board) {
		return nil
	}

	if len(board[col]) > 0 && board[col][0] != 0 {
		return errCollumnFull
	}

	for row := 0; row < len(board[col]); row++ {
		if board[col][row] != 0 {
			board[col][row-1] = val
			return nil
		}
	}

	board[col][len(board[col])-1] = val
	return nil
}

func utility(board [][]byte, val byte, col int) float64 {
	b := clone(board)
	current := score(board, val)
	drop(b, val, col)
	return score(b, val) - current
}

func display(board [][]byte) {
	for row := 0; row < len(board[len(board)-1]); row++ {
		fmt.Printf("|")
		for col := 0; col < len(board); col++ {
			if board[col][row] != 0 {
				fmt.Printf(" %c |", board[col][row])
			} else {
				fmt.Print("   |")
			}
		}
		fmt.Println()
	}

	// for range board {
	// 	fmt.Print("----")
	// }

	// fmt.Println()
	// fmt.Print("|")
	for i := range board {
		fmt.Printf(" %2d ", i)
	}
	fmt.Println()
}

func clone(board [][]byte) [][]byte {
	cloned := newBoard(len(board), len(board[0]))

	for col := range board {
		for row := range board[col] {
			cloned[row][col] = board[row][col]
		}
	}

	return cloned
}

func score(board [][]byte, val byte) float64 {
	if _, done := terminal(board, val); done {
		return winningScore
	}

	if findPatternWithMask(board, val, mask4Vertical) ||
		findPatternWithMask(board, val, mask4Horzontal) ||
		findPatternWithMask(board, val, mask4ForwardDiagnal) ||
		findPatternWithMask(board, val, mask4BackwardsDiagnal) {
		return 1
	}
	return 0

}

func terminal(board [][]byte, vals ...byte) (byte, bool) {
	for _, val := range vals {
		if findPatternWithMask(board, val, mask4Vertical) ||
			findPatternWithMask(board, val, mask4Horzontal) ||
			findPatternWithMask(board, val, mask4ForwardDiagnal) ||
			findPatternWithMask(board, val, mask4BackwardsDiagnal) {
			return val, true
		}
	}
	return byte(0), full(board)
}

func full(board [][]byte) bool {
	hasPossibleNextMove := false
	for boardCol := 0; boardCol < len(board) && !hasPossibleNextMove; boardCol++ {
		if board[boardCol][0] == byte(0) {
			hasPossibleNextMove = true
		}
	}
	return !hasPossibleNextMove
}

// func check(board [][]byte, val byte, c, r, sizeCol, sizeRow, incrementCol, incrementRow int) int {
// 	count := 0
// 	fmt.Println(c, r)
//
// 	if c+sizeCol >= len(board) || r+sizeRow > len(board[0]) {
// 		return count
// 	}
//
// 	for col := c; col < sizeCol+c; col += incrementCol {
// 		for row := r; row < sizeRow+r; row += incrementRow {
// 			if board[col][row] == val {
// 				count++
// 			} else if board[col][row] == byte(0) {
// 				// empty spot
// 			} else {
// 				return 0
// 			}
// 		}
// 	}
// 	return count
// }
//
