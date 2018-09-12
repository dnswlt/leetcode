package leetcode

func diff1(v, w string) bool {
	d := 0
	for i := 0; i < len(v); i++ {
		if v[i] != w[i] {
			d++
		}
	}
	return d == 1
}

type searchState struct {
	i     int
	depth int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	g := map[int][]int{}
	endFound := false
	wordList = append(wordList, beginWord)
	for i := 0; i < len(wordList); i++ {
		endFound = endFound || wordList[i] == endWord
		for j := i + 1; j < len(wordList); j++ {
			if diff1(wordList[i], wordList[j]) {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}
	if !endFound {
		return 0
	}
	q := []searchState{searchState{len(wordList) - 1, 0}}
	seen := make(map[int]struct{})
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		if wordList[s.i] == endWord {
			return s.depth + 1
		}

		seen[s.i] = struct{}{}
		succs := g[s.i]
		for i := 0; i < len(succs); i++ {
			if _, ok := seen[succs[i]]; !ok {
				q = append(q, searchState{succs[i], s.depth + 1})
			}
		}
	}
	return 0
}
