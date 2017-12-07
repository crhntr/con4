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

func drop(board [][]byte, val byte, col int) ([][]byte, error) {
	if col < 0 || col >= len(board) {
		return board, nil
	}

	if len(board[col]) > 0 && board[col][0] != 0 {
		return board, errCollumnFull
	}

	for row := 0; row < len(board[col]); row++ {
		if board[col][row] != 0 {
			board[col][row-1] = val
			return board, nil
		}
	}

	board[col][len(board[col])-1] = val
	return board, nil
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

	return findPatternWithMask(board, val, maskVertical, 0.25) +
		findPatternWithMask(board, val, maskHorzontal, 0.25) +
		findPatternWithMask(board, val, maskForwardDiagnal, 0.25) +
		findPatternWithMask(board, val, maskBackwardsDiagnal, 0.25)
}

func findPatternWithMask(board [][]byte, val byte, mask [][]bool, score float64) float64 {
	acumulatedScore := 0.0

	incrementCol := 1
	incrementRow := 1
	if len(mask) == len(mask[0]) {
		incrementCol = len(mask)
		incrementRow = len(mask[0])
	}

	for boardCol := len(board) - 1; boardCol >= 0; boardCol -= incrementCol {
		for boardRow := len(board) - 1; boardRow >= 0; boardRow -= incrementRow {

			for maskCol := len(mask) - 1; maskCol >= 0 && boardCol+maskCol >= 0 && boardCol+maskCol < len(board); maskCol-- {
				for maskRow := len(mask[maskCol]) - 1; maskRow >= 0 && boardRow+maskRow >= 0 && boardRow+maskRow < len(board[boardCol+maskCol]); maskRow-- {

					if mask[maskCol][maskRow] && board[boardCol+maskCol][boardRow+maskRow] == val {
						acumulatedScore += score
					}
				}
			}
		}
	}

	return acumulatedScore
}

func terminal(board [][]byte, vals ...byte) (byte, bool) {
	for _, val := range vals {

		for _, mask := range [][][]bool{
			maskVertical,
			maskHorzontal,
			maskForwardDiagnal,
			maskBackwardsDiagnal} {

			incrementCol := 1
			incrementRow := 1
			if len(mask) == len(mask[0]) {
				incrementCol = len(mask)
				incrementRow = len(mask[0])
			}

			for boardCol := len(board) - 1; boardCol >= 0; boardCol -= incrementCol {
				boardRow := len(board) - 1
			nextFrame:

				for ; boardRow >= 0; boardRow -= incrementRow {
					firstMatchFound := false

					for maskCol := len(mask) - 1; maskCol >= 0 && boardCol+maskCol >= 0 && boardCol+maskCol < len(board); maskCol-- {
						for maskRow := len(mask[maskCol]) - 1; maskRow >= 0 && boardRow+maskRow >= 0 && boardRow+maskRow < len(board[boardCol+maskCol]); maskRow-- {

							if mask[maskCol][maskRow] {

								if board[boardCol+maskCol][boardRow+maskRow] == val {
									if firstMatchFound && maskCol == 0 && maskRow == 0 {
										return val, true
									}
									firstMatchFound = true

								} else {
									continue nextFrame
								}

							}
						}
					}
				}
			}
		}

	}
	return byte(0), false
}

var maskVertical = [][]bool{
	{true, true, true, true},
}

var maskHorzontal = [][]bool{
	{true},
	{true},
	{true},
	{true},
}

var maskForwardDiagnal = [][]bool{
	{true, false, false, false},
	{false, true, false, false},
	{false, false, true, false},
	{false, false, false, true},
}

var maskBackwardsDiagnal = [][]bool{
	{false, false, false, true},
	{false, false, true, false},
	{false, true, false, false},
	{true, false, false, false},
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
