package main

type Moves []moveScore

type moveScore struct {
	Move  uint64
	Score int
}

func (list *Moves) Add(move uint64, score int) { (*list) = append((*list), moveScore{move, score}) }
func (list *Moves) Len() int                   { return len(*list) }
func (list *Moves) Swap(i, j int)              { (*list)[i], (*list)[j] = (*list)[j], (*list)[i] }
func (list *Moves) Less(i, j int) bool         { return (*list)[i].Score < (*list)[j].Score }
