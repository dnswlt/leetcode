package leetcode

import "sort"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	best := map[int]int{}
	for a := 1; a <= amount; a++ {
		findBest(coins, best, a)
	}
	v, ok := best[amount]
	if ok {
		return v
	}
	return -1
}

func findBest(coins []int, best map[int]int, amount int) {
	for i := 0; i < len(coins); i++ {
		rem := amount - coins[i]
		if rem < 0 {
			return
		} else if rem == 0 {
			best[amount] = 1
		} else {
			v, ok := best[rem]
			if ok {
				curr, ok := best[amount]
				if !ok || curr > v+1 {
					best[amount] = v + 1
				}
			}
		}
	}
}
