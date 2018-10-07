package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	lp := longestPalindrome("meinotto.de")
	if lp != "otto" {
		t.Error("Funny longest palindrome", lp)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	lp := longestPalindrome("malayalam")
	if lp != "malayalam" {
		t.Error("Funny longest palindrome", lp)
	}
}

func TestLongestPalindrome3(t *testing.T) {
	lp := longestPalindrome("aacdefcaa")
	if lp != "aa" {
		t.Error("Funny longest palindrome", lp)
	}
}

func TestLongestPalindrome4(t *testing.T) {
	lp := longestPalindrome("aa")
	if lp != "aa" {
		t.Error("Funny longest palindrome", lp)
	}
}

func TestLongestPalindrome5(t *testing.T) {
	lp := longestPalindrome("abc")
	if lp != "a" {
		t.Error("Funny longest palindrome", lp)
	}
}
