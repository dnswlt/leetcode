package leetcode

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	combs := letterCombinations("23")
	if fmt.Sprint(combs) != "[ad ae af bd be bf cd ce cf]" {
		t.Error("Not such a nice comb", combs)
	}
}
