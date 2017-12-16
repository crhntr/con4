package main

import (
	"errors"
	"fmt"
)

var (
	errCollumnFull = errors.New("collumn full")
)

type ByteBoard [][]byte

func NewByteBoard(cols, rows int) ByteBoard {
	board := make(ByteBoard, cols)
	for col := range board {
		board[col] = make([]byte, rows)
	}
	return board
}

func (board ByteBoard) Drop(val byte, col int) error {
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

func (board ByteBoard) Utility(val byte, col int) float64 {
	b := board.Clone()
	currentScore := board.Score(val)
	b.Drop(val, col)
	return b.Score(val) - currentScore
}

func (board ByteBoard) String() string {
	str := ""
	for row := 0; row < len(board[len(board)-1]); row++ {
		str += fmt.Sprintf("|")
		for col := 0; col < len(board); col++ {
			if board[col][row] != 0 {
				str += fmt.Sprintf(" %c |", board[col][row])
			} else {
				str += fmt.Sprint("   |")
			}
		}
		str += fmt.Sprintln()
	}

	for i := range board {
		str += fmt.Sprintf(" %2d ", i)
	}
	return str + "\n"
}

func (board ByteBoard) Clone() ByteBoard {
	cloned := NewByteBoard(len(board), len(board[0]))

	for col := range board {
		for row := range board[col] {
			cloned[row][col] = board[row][col]
		}
	}

	return cloned
}

func (board ByteBoard) Score(val byte) float64 {
	if _, done := board.IsTerminal(val); done {
		return winningScore
	}

	if board.findPatternWithMask(val, mask4Vertical) ||
		board.findPatternWithMask(val, mask4Horzontal) ||
		board.findPatternWithMask(val, mask4ForwardDiagnal) ||
		board.findPatternWithMask(val, mask4BackwardsDiagnal) {
		return 1
	}
	return 0

}

func (board ByteBoard) IsTerminal(vals ...byte) (byte, bool) {
	for _, val := range vals {
		if board.findPatternWithMask(val, mask4Vertical) ||
			board.findPatternWithMask(val, mask4Horzontal) ||
			board.findPatternWithMask(val, mask4ForwardDiagnal) ||
			board.findPatternWithMask(val, mask4BackwardsDiagnal) {
			return val, true
		}
	}
	return byte(0), board.IsFull()
}

func (board ByteBoard) IsFull() bool {
	hasPossibleNextMove := false
	for boardCol := 0; boardCol < len(board) && !hasPossibleNextMove; boardCol++ {
		if board[boardCol][0] == byte(0) {
			hasPossibleNextMove = true
		}
	}
	return !hasPossibleNextMove
}
