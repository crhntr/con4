package main

import (
	"fmt"
	"math"
)

const (
	maxDepth     = 12
	winningScore = 1000
)

func main() {
	blueChip, redChip := byte('O'), byte('X')

	board := NewByteBoard(8, 8)
	machineTurn := false
	for {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

		if winner, done := board.IsTerminal(blueChip, redChip); done {
			fmt.Println(board)
			fmt.Printf("%c has won!\n", winner)
			break
		}

		var err error
		if machineTurn {
			fmt.Print("machine thinking...")
			machineSelection := solvenegamax(board, redChip, blueChip)
			err = board.Drop(redChip, machineSelection)
			fmt.Printf("and chose %d\n", machineSelection)
		} else {
			fmt.Println(board)

			fmt.Println("select a column number: ")
			var userSelection int
			fmt.Scanf("%d", &userSelection)

			err = board.Drop(blueChip, userSelection)
		}
		if err != nil {
			fmt.Println("column selected did not work")
		}
		machineTurn = !machineTurn
	}
}

func solvenegamax(board ByteBoard, val1, val2 byte) int {
	scores := make([]int, len(board))

	max := len(board) * len(board[0])

	b := board.Clone()

	for i := range scores {
		scores[i] = negamax(b, val1, val2, max/2, -max/2, 0)
	}

	fmt.Println(scores)
	return maxInt(scores)
}

func negamax(board ByteBoard, val1, val2 byte, alpha, beta, moves int) int {
	if board.IsFull() {
		return 0
	}

	moves++

	max := len(board)*len(board[0]) - moves/2

	for col := 0; col < len(board); col++ {
		b := board.Clone()
		err := b.Drop(val1, col)
		if _, term := b.IsTerminal(val1); err == nil && term {
			return max
		}
	}

	if beta > max {
		beta = max
		if alpha >= beta {
			return beta
		}
	}

	for col := 0; col < len(board); col++ {
		b := board.Clone()
		if err := board.Drop(val1, col); err == nil {
			score := -negamax(b, val2, val1, -beta, -alpha, moves)
			if score >= beta {
				return score
			}
			if score > alpha {
				alpha = score
			}
		}
	}
	return alpha
}

func maxInt(ints []int) int {
	max, maxIndex := math.MinInt64, 0
	for i := range ints {
		if ints[i] > max {
			max, maxIndex = ints[i], i
		}
	}
	return maxIndex
}
