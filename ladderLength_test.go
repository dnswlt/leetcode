package leetcode

import "testing"

func TestLadderLength(t *testing.T) {
	l := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	if l != 5 {
		t.Error("Not the right length", l)
	}
}

func TestLadderLength2(t *testing.T) {
	l := ladderLength("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"})
	if l != 5 {
		t.Error("Not the right length", l)
	}
}
