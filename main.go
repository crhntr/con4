package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	maxDepth     = 12
	winningScore = 1000
)

func main() {
	blueChip, redChip := byte('O'), byte('X')

	board := newBoard(8, 8)
	rounds := 0

	for {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n")

		if winner, done := terminal(board, blueChip, redChip); done {
			fmt.Printf("%c has won!\n", winner)
			break
		}

		fmt.Print("machine thinking...")
		machineSelection := rand.Intn(len(board))
		if rounds > 2 {
			machineSelection = MinimaxDecision(board, blueChip, redChip, maxDepth)
		}
		drop(board, redChip, machineSelection)
		fmt.Printf("and chose %d\n", machineSelection)

		fmt.Println()
		display(board)

		fmt.Println("select a column number: ")
		var userSelection int
		fmt.Scanf("%d", &userSelection)

		drop(board, blueChip, userSelection)
		display(board)
		fmt.Println()
	}
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
