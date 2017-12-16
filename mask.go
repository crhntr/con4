package main

var mask4Vertical = [][]bool{
	{true, true, true, true},
}

var mask4Horzontal = [][]bool{
	{true},
	{true},
	{true},
	{true},
}

var mask4ForwardDiagnal = [][]bool{
	{true, false, false, false},
	{false, true, false, false},
	{false, false, true, false},
	{false, false, false, true},
}

var mask4BackwardsDiagnal = [][]bool{
	{false, false, false, true},
	{false, false, true, false},
	{false, true, false, false},
	{true, false, false, false},
}

func findPatternWithMask(board [][]byte, val byte, mask [][]bool) bool {
	for boardCol := 0; boardCol <= len(board)-len(mask); boardCol++ {
		for boardRow := 0; boardRow <= len(board)-len(mask[0]); boardRow++ {
			accross := 0

			for maskCol := 0; maskCol < len(mask); maskCol++ {
				for maskRow := 0; maskRow < len(mask[maskCol]); maskRow++ {

					if mask[maskCol][maskRow] && board[boardCol+maskCol][boardRow+maskRow] == val {
						accross++
					}

				}
			}

			if accross > 3 {
				return true
			}
		}
	}

	return false
}
