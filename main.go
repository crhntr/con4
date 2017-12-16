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
			fmt.Printf("%c has won!\n", winner)
			break
		}

		var err error
		if machineTurn {
			fmt.Print("machine thinking...")
			_, machineSelection := NegaMax(board, redChip, blueChip)
			err = drop(board, redChip, machineSelection)
			fmt.Printf("and chose %d\n", machineSelection)
		} else {
			fmt.Println()
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

var depth = 0

func NegaMax(board [][]byte, val1, val2 byte) (score, column int) {
	if full(board) {
		return 0, 0
	}

	for col := 0; col < len(board); col++ {
		b := clone(board)
		err := drop(b, val1, col)
		if _, term := terminal(b, val1); err == nil && term {
			score, column = 1, col
			return score, column
		}
	}

	score, column = math.MinInt64, len(board)/2

	for col := 0; col < len(board); col++ {
		b := clone(board)
		if err := drop(b, val1, col); err == nil {
			s, _ := NegaMax(b, val2, val1)
			s = -s
			if s > score {
				score, column = s, col
			}
		}
	}
	return score, column
}

func MinimaxDecision(board [][]byte, minVal, maxVal byte, depth int) int {
	b := clone(board)

	alpha, beta := math.Inf(-1), math.Inf(1)

	bestCol, maxUtility := 0, -math.Inf(-1)

	// scores := make([]float64, len(b))
	for c := 0; c < len(b); c++ {
		drop(board, maxVal, c)
		v := minValue(board, alpha, beta, minVal, maxVal, maxDepth)
		// scores = append(scores, v)

		if v > maxUtility {
			bestCol, maxUtility = c, v

			if alpha > v {
				alpha = v
			}
		}
	}

	return bestCol
}

func maxValue(board [][]byte, alpha, beta float64, minVal, maxVal byte, depth int) float64 {

	depth--
	if _, done := terminal(board, minVal, maxVal); done || depth <= 0 {
		return score(board, maxVal) - score(board, minVal)
	}

	v := math.Inf(1)

	for c := 0; c < len(board); c++ {
		b := clone(board)
		drop(board, maxVal, c)

		utilVal := minValue(b, alpha, beta, minVal, maxVal, depth)
		if utilVal > v {
			v = utilVal
		}

		if v <= beta {
			return v
		}

		if v > alpha {
			alpha = v
		}
	}

	return v
}

func minValue(board [][]byte, alpha, beta float64, minVal, maxVal byte, depth int) float64 {

	depth--
	if _, done := terminal(board, minVal, maxVal); done || depth <= 0 {
		return score(board, maxVal) - score(board, minVal)
	}

	v := math.Inf(-1)

	for c := 0; c < len(board); c++ {
		b := clone(board)
		drop(board, minVal, c)

		utilVal := minValue(b, alpha, beta, minVal, maxVal, depth)
		if utilVal < v {
			v = utilVal
		}

		if v >= alpha {
			return v
		}

		if v < beta {
			beta = v
		}
	}

	return v
}
