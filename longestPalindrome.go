package leetcode

func longestRevertedSubstring(s string) string {
	// 1st attempt at finding the longest palindrome:
	// Unfortunately, finds the longest substring that exists both in forward and reverse direction in s :(
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	maxI, maxLen := -1, 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == s[n-j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxI, maxLen = i, dp[i][j]
				}
			}
		}
	}
	if maxI == -1 {
		return ""
	}
	return s[maxI-maxLen : maxI]
}

func isPalindrome(s string, i, j int, dp [][]int8) int8 {
	if i > j {
		return 1
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	var b int8 = -1
	if s[i] == s[j] {
		b = isPalindrome(s, i+1, j-1, dp)
	}
	dp[i][j] = b
	return b
}

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]int8, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int8, n)
	}
	maxI, maxLen := -1, 0
	for i := 0; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if isPalindrome(s, i, j, dp) == 1 {
				if j-i+1 > maxLen {
					maxI, maxLen = i, j-i+1
					break
				}
			}
		}
	}
	if maxI == -1 {
		return ""
	}
	return s[maxI : maxI+maxLen]
}
