package leetcode

import "testing"

func TestCoinChange250(t *testing.T) {
	coins := []int{5, 10, 20, 50, 100, 200, 500}
	n := coinChange(coins, 250)
	if n != 2 {
		t.Error("Not the right number of coins", n)
	}
}

func TestCoinChange60(t *testing.T) {
	coins := []int{5, 10, 20, 50, 100, 200, 500}
	n := coinChange(coins, 60)
	if n != 2 {
		t.Error("Not the right number of coins", n)
	}
}

func TestCoinChange264(t *testing.T) {
	coins := []int{474, 83, 404, 3}
	n := coinChange(coins, 264)
	if n != 8 {
		t.Error("Not the right number of coins", n)
	}
}
