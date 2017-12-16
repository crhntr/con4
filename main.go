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

	board := newBoard(8, 8)
	machineTurn := false
	for {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

		if winner, done := terminal(board, blueChip, redChip); done {
			display(board)
			fmt.Printf("%c has won!\n", winner)
			break
		}

		var err error
		if machineTurn {
			fmt.Print("machine thinking...")
			machineSelection := solvenegamax(board, redChip, blueChip)
			err = drop(board, redChip, machineSelection)
			fmt.Printf("and chose %d\n", machineSelection)
		} else {
			display(board)

			fmt.Println("select a column number: ")
			var userSelection int
			fmt.Scanf("%d", &userSelection)

			err = drop(board, blueChip, userSelection)
		}
		if err != nil {
			fmt.Println("column selected did not work")
		}
		machineTurn = !machineTurn
	}
}

func solvenegamax(board [][]byte, val1, val2 byte) int {
	scores := make([]int, len(board))

	max := len(board) * len(board[0])

	for i := range scores {
		b := clone(board)
		if err := drop(b, val1, i); err == nil {
			scores[i] = negamax(b, val1, val2, max-1, -max+1, 0)
		} else {
			scores[i] = -max
		}
	}

	fmt.Println(scores)
	return maxInt(scores)
}

func negamax(board [][]byte, val1, val2 byte, alpha, beta, moves int) int {
	if full(board) {
		return 0
	}

	moves++

	max := len(board)*len(board[0]) - moves/2

	for col := 0; col < len(board); col++ {
		b := clone(board)
		err := drop(b, val1, col)
		if _, term := terminal(b, val1); err == nil && term {
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
		b := clone(board)
		if err := drop(b, val1, col); err == nil {
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
