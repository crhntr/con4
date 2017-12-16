package main

import (
	"fmt"
	"math/bits"
)

const (
	Width    = 7 // width of the board
	Height   = 6 // height of the board
	MinScore = -(Width*Height)/2 + 3
	MaxScore = (Width*Height+1)/2 - 3
)

var (
	bottomMask = bottom(Width, Height)
	boardMask  = bottomMask * ((1 << Height) - 1)
)

func init() {
	if Width >= 10 {
		panic("width must be less than 10")
	}
	if Width*(Height+1) > 64 {
		panic("board does not fit into a 64 bit bitboard")
	}
}

func ColumnMask(col uint64) uint64 {
	return ((1 << Height) - 1) << col * (Height + 1)
}

func bottom(width uint64, height uint64) uint64 {
	if width == 0 {
		return 0
	}
	return bottom(width-1, height) | 1<<((width-1)*(height+1))
}

func topMaskColumn(col uint64) uint64 {
	return uint64(1) << ((Height - 1) + col*(Height+1))
}

// return a bitmask containg a single 1 corresponding to the bottom cell of a given column
func bottomMaskColumn(col uint64) uint64 {
	return uint64(1) << col * (Height + 1)
}

func getBit(n uint64, pos int) uint64 {
	val := n & (1 << uint(pos))
	if val > 0 {
		return 1
	}
	return 0
}

type Board struct {
	currentBoard uint64 // bitmap of the currentPlayer's tokens
	mask         uint64 // bitmap of all the already played spots
	Moves        int    // number of moves played
}

func (board Board) Clone() Board {
	return Board{
		currentBoard: board.currentBoard,
		mask:         board.mask,
		Moves:        board.Moves,
	}
}

func (board Board) String() string {
	opChar, cpChar := "x ", "o "
	if ((board.Moves % 2) + 1) == 1 {
		opChar, cpChar = cpChar, opChar
	}

	buf := ""

	for r := 0; r < Height; r++ {
		rBuf := ""

		for c := 0; c < Width*(Height+1); c += Height + 1 {
			current := getBit(board.currentBoard, c+r)
			mask := getBit(board.mask, c+r)

			if current == 1 && mask == 1 {
				rBuf += opChar
			} else if current == 0 && mask == 1 {
				rBuf += cpChar
			} else {
				rBuf += ". "
			}
		}

		buf = buf + "\n" + rBuf
	}

	return buf
}

func (board *Board) Play(move uint64) {
	board.currentBoard ^= board.mask
	board.mask |= move
	board.Moves++
}

func (board Board) PossibleNonLosingMoves() uint64 {
	possibleMask := board.possible()
	opponentWin := board.opponentWinningBoard()
	forcedMoves := possibleMask & opponentWin

	if forcedMoves != 0 {
		if forcedMoves&(forcedMoves-1) != 0 {
			return 0
		}
		opponentWin = forcedMoves
	}
	return possibleMask & ^(opponentWin >> 1)
}

func (board Board) CanWinNext() bool {
	return board.winningBoard()&board.possible() != 0
}

func (board Board) MoveScore(move uint64) int {
	return bits.OnesCount64((ComputeWinningBoard(board.currentBoard|move, board.mask)))
}

func (board *Board) PlayColumn(column int) {
	board.Play((board.mask + bottomMaskColumn(uint64(column))) & ColumnMask(uint64(column)))
}

func (board *Board) PlayAll(moveColumns ...int) int {
	for i, col := range moveColumns {
		if !board.canPlay(uint64(col)) {
			panic(fmt.Errorf("board cannot play %d", col))
		}

		if col < 0 || col > Width || board.IsWinningMove(col) {
			return i
		}
		board.PlayColumn(col)
	}
	return len(moveColumns)
}

func (board *Board) PlayString(plays string) int {
	moves := []int{}
	for i, _ := range plays {
		moves = append(moves, int(plays[i])-int('0'))
	}
	return board.PlayAll(moves...)
}

func (board *Board) Key() uint64 {
	return board.currentBoard | board.mask
}

func (board Board) IsWinningMove(col int) bool {
	return (board.winningBoard() & board.possible() & ColumnMask(uint64(col))) != 0
}

func (board Board) canPlay(col uint64) bool {
	return (board.mask & topMaskColumn(col)) == 0
}

func (board Board) winningBoard() uint64 {
	return ComputeWinningBoard(board.currentBoard, board.mask)
}

func (board Board) opponentWinningBoard() uint64 {
	return ComputeWinningBoard(board.currentBoard^board.mask, board.mask)
}

func (board Board) possible() uint64 {
	return (board.mask | bottomMask) & boardMask
}

func ComputeWinningBoard(board uint64, mask uint64) uint64 {
	// vertical
	r := (board << 1) & (board << 2) & (board << 3)

	// horizontal
	p := (board << (Height + 1)) & (board << 2 * (Height + 1))
	r |= p & (board << 3 * (Height + 1))
	r |= p & (board >> (Height + 1))
	p = (board >> (Height + 1)) & (board >> 2 * (Height + 1))
	r |= p & (board << (Height + 1))
	r |= p & (board >> 3 * (Height + 1))

	// diagonal 1
	p = (board << Height) & (board << 2 * Height)
	r |= p & (board << 3 * Height)
	r |= p & (board >> Height)
	p = (board >> Height) & (board >> 2 * Height)
	r |= p & (board << Height)
	r |= p & (board >> 3 * Height)

	// diagonal 2
	p = (board << (Height + 2)) & (board << 2 * (Height + 2))
	r |= p & (board << 3 * (Height + 2))
	r |= p & (board >> (Height + 2))
	p = (board >> (Height + 2)) & (board >> 2 * (Height + 2))
	r |= p & (board << (Height + 2))
	r |= p & (board >> 3 * (Height + 2))

	return r & (boardMask ^ mask)
}
