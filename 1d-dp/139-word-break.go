package ddp

import "math"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	dict := make(map[string]bool)

	for _, w := range wordDict {
		dict[w] = true
	}

	minLen, maxLen := math.MaxInt, 0
	for w := range dict {
		if len(w) < minLen {
			minLen = len(w)
		}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		start := i - maxLen
		if start < 0 {
			start = 0
		}
		for j := start; j <= i-minLen; j++ {
			if !dp[j] {
				continue
			}

			w := s[j:i]
			if _, ok := dict[w]; ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
