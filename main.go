package main

import "fmt"

var (
	nodeCount   = 0
	columnOrder [Width]uint64
)

func init() {
	for i := 0; i < Width; i++ {
		columnOrder[i] = uint64(Width/2 + (1-2*(i%2))*(i+1)/2)
	}
}

func main() {
	board := Board{}

	board.PlayColumn(0)
	board.PlayColumn(1)
	board.PlayColumn(0)
	board.PlayColumn(1)
	board.PlayColumn(0)
	board.PlayColumn(1)
	board.PlayColumn(0)
	board.PlayColumn(1)

	fmt.Println(board)

	fmt.Printf("%064b\n%064b\nc", board.currentBoard, board.mask)
}

func solve(board Board, weak bool) int {
	if board.CanWinNext() {
		return (Width*Height + 1 - board.Moves) / 2
	} // check if win in one move as the Negamax function does not support this case.

	min := -(Width*Height - board.Moves) / 2
	max := (Width*Height + 1 - board.Moves) / 2
	if weak {
		min = -1
		max = 1
	}

	for min < max { // iteratively narrow the min-max exploration window
		med := min + (max-min)/2
		if med <= 0 && min/2 < med {
			med = min / 2
		} else if med >= 0 && max/2 > med {
			med = max / 2
		}

		r := negamax(board, med, med+1) // use a null depth window to know if the actual score is greater or smaller than med
		if r <= med {
			max = r
		} else {
			min = r
		}
	}
	return min
}

func negamax(board Board, alpha, beta int) int {
	// assert(alpha < beta);
	// assert(!board.canWinNext());

	nodeCount++ // increment counter of explored nodes

	possible := board.PossibleNonLosingMoves()
	if possible == 0 {
		return -(Width*Height - board.Moves) / 2 // opponent wins next move
	}

	if board.Moves > Width*Height-2 {
		return 0 // tie
	}

	min := -(Width*Height - 2 - board.Moves) / 2 // lower bound of score as opponent cannot win next move
	if alpha < min {
		alpha = min // there is no need to keep beta above our max possible score.
		if alpha >= beta {
			return alpha
		} // prune the exploration if the [alpha;beta] window is empty.
	}

	max := (Width*Height - 1 - board.Moves) / 2 // upper bound of our score as we cannot win immediately
	if beta > max {
		beta = max // there is no need to keep beta above our max possible score.
		if alpha >= beta {
			return beta
		} // prune the exploration if the [alpha;beta] window is empty.
	}

	stateMap := map[uint64]int{}

	key := board.Key()

	if val, found := stateMap[key]; found {
		if val > MaxScore-MinScore+1 { // we have an lower bound

			min = val + 2*MinScore - MaxScore - 2

			if alpha < min {
				alpha = min // there is no need to keep beta above our max possible score.
				if alpha >= beta {
					return alpha
				} // prune the exploration if the [alpha;beta] window is empty.
			}
		} else { // we have an upper bound

			max = val + MinScore - 1

			if beta > max {
				beta = max // there is no need to keep beta above our max possible score.
				if alpha >= beta {
					return beta
				} // prune the exploration if the [alpha;beta] window is empty.
			}
		}
	}

	moves := Moves{}

	for i := Width; i >= 0; i-- {
		if move := possible & ColumnMask(columnOrder[i]); move > 0 {
			moves.Add(move, board.MoveScore(move))
		}
	}

	for _, move := range moves {
		nextBoard := board.Clone()

		nextBoard.Play(move.Move)

		score := -negamax(nextBoard, -beta, -alpha)

		if score >= beta {
			moves.Add(board.Key(), score+MaxScore-2*MinScore+2)
			return score
		}
		if score > alpha {
			alpha = score
		}
	}

	stateMap[key] = alpha - MinScore + 1 // save the upper bound of the position
	return alpha
}
