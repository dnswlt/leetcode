package leetcode

import "sort"

type job struct {
	difficulty int
	profit     int
}

type jobs []job

func (s jobs) Len() int {
	return len(s)
}

func (s jobs) Less(i, j int) bool {
	return s[i].difficulty < s[j].difficulty
}

func (s jobs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	l := len(difficulty)
	total := 0
	jobs := make(jobs, l)
	for i := 0; i < l; i++ {
		jobs[i] = job{difficulty[i], profit[i]}
	}
	sort.Sort(jobs)
	for i := 1; i < l; i++ {
		if jobs[i-1].profit > jobs[i].profit {
			jobs[i].profit = jobs[i-1].profit
		}
	}
	for _, w := range worker {
		i := sort.Search(l, func(j int) bool { return jobs[j].difficulty > w })
		if i > 0 {
			total += jobs[i-1].profit
		}
	}
	return total
}
