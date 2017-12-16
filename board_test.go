package main

import "testing"

func TestBoardTerminalDown(t *testing.T) {
	b := NewByteBoard(8, 8)
	val := byte('o')
	b.Drop(val, 0)
	b.Drop(val, 0)
	b.Drop(val, 0)
	if _, term := b.IsTerminal(val); term {
		t.Error("should not be terminal")
	}
	b.Drop(val, 0)

	if _, term := b.IsTerminal(val); !term {
		t.Error("should be terminal")
	}
}

func TestBoardTerminalAcross(t *testing.T) {
	b := NewByteBoard(8, 8)
	val := byte('o')
	b.Drop(val, 0)
	b.Drop(val, 1)
	b.Drop(val, 2)
	if _, term := b.IsTerminal(val); term {
		t.Error("should not be terminal")
	}
	b.Drop(val, 3)
	b.Drop(val, 4)

	if _, term := b.IsTerminal(val); !term {
		t.Error("should be terminal")
	}
}

func TestBoardMatchesMask(t *testing.T) {
	b := NewByteBoard(8, 8)
	val1, val2 := byte('o'), byte('x')
	b.Drop(val1, 0)

	b.Drop(val2, 1)
	b.Drop(val1, 1)

	b.Drop(val2, 2)
	b.Drop(val2, 2)
	b.Drop(val1, 2)

	b.Drop(val2, 3)
	b.Drop(val2, 3)
	b.Drop(val2, 3)
	b.Drop(val1, 3)

	if !b.findPatternWithMask(val1, mask4BackwardsDiagnal) {
		t.Log(b)
		t.Error("should not be terminal")
	}

}
