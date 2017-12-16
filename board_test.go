package main

import "testing"

func TestBoardTerminalDown(t *testing.T) {
	b := newBoard(8, 8)
	val := byte('o')
	drop(b, val, 0)
	drop(b, val, 0)
	drop(b, val, 0)
	if _, term := terminal(b, val); term {
		t.Error("should not be terminal")
	}
	drop(b, val, 0)

	if _, term := terminal(b, val); !term {
		t.Error("should be terminal")
	}
}

func TestBoardTerminalAcross(t *testing.T) {
	b := newBoard(8, 8)
	val := byte('o')
	drop(b, val, 0)
	drop(b, val, 1)
	drop(b, val, 2)
	if _, term := terminal(b, val); term {
		t.Error("should not be terminal")
	}
	drop(b, val, 3)
	drop(b, val, 4)

	if _, term := terminal(b, val); !term {
		t.Error("should be terminal")
	}
}

func TestBoardMatchesMask(t *testing.T) {
	b := newBoard(8, 8)
	val1, val2 := byte('o'), byte('x')
	drop(b, val1, 0)

	drop(b, val2, 1)
	drop(b, val1, 1)

	drop(b, val2, 2)
	drop(b, val2, 2)
	drop(b, val1, 2)

	drop(b, val2, 3)
	drop(b, val2, 3)
	drop(b, val2, 3)
	drop(b, val1, 3)

	if !findPatternWithMask(b, val1, mask4BackwardsDiagnal) {
		t.Log(b)
		t.Error("should not be terminal")
	}

}
