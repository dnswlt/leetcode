package leetcode

import (
	"strings"
	"testing"
)

func splitBoard(s string) [][]byte {
	lines := strings.Split(s, "\n")
	r := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = []byte(lines[i])
	}
	return r
}

func stringBoard(board [][]byte) string {
	r := make([]string, len(board))
	for i := 0; i < len(r); i++ {
		r[i] = string(board[i])
	}
	return strings.Join(r, "\n")
}

func TestSurroundedRegions(t *testing.T) {
	b := splitBoard(`XXXX
XOOX
XXOX
XOXX`)
	surroundedRegions(b)
	if stringBoard(b) != `XXXX
XXXX
XXXX
XOXX` {
		t.Error(stringBoard(b))
	}
}
