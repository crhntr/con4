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
			fmt.Printf("%c has won!\n", winner)
			break
		}

		var err error
		if machineTurn {
			fmt.Print("machine thinking...")
			_, machineSelection := NegaMax(board, redChip, blueChip)
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

var depth = 0

func NegaMax(board ByteBoard, val1, val2 byte) (score, column int) {
	if board.IsFull() {
		return 0, 0
	}

	for col := 0; col < len(board); col++ {
		b := board.Clone()
		err := b.Drop(val1, col)
		if _, term := b.IsTerminal(val1); err == nil && term {
			score, column = 1, col
			return score, column
		}
	}

	score, column = math.MinInt64, len(board)/2

	for col := 0; col < len(board); col++ {
		b := board.Clone()
		if err := board.Drop(val1, col); err == nil {
			s, _ := NegaMax(b, val2, val1)
			s = -s
			if s > score {
				score, column = s, col
			}
		}
	}
	return score, column
}

func MinimaxDecision(board ByteBoard, minVal, maxVal byte, depth int) int {
	b := board.Clone()

	alpha, beta := math.Inf(-1), math.Inf(1)

	bestCol, maxUtility := 0, -math.Inf(-1)

	// scores := make([]float64, len(b))
	for c := 0; c < len(b); c++ {
		board.Drop(maxVal, c)
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

func maxValue(board ByteBoard, alpha, beta float64, minVal, maxVal byte, depth int) float64 {

	depth--
	if _, done := board.IsTerminal(minVal, maxVal); done || depth <= 0 {
		return board.Score(maxVal) - board.Score(minVal)
	}

	v := math.Inf(1)

	for c := 0; c < len(board); c++ {
		b := board.Clone()
		board.Drop(maxVal, c)

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

func minValue(board ByteBoard, alpha, beta float64, minVal, maxVal byte, depth int) float64 {

	depth--
	if _, done := board.IsTerminal(minVal, maxVal); done || depth <= 0 {
		return board.Score(maxVal) - board.Score(minVal)
	}

	v := math.Inf(-1)

	for c := 0; c < len(board); c++ {
		b := board.Clone()
		board.Drop(minVal, c)

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
