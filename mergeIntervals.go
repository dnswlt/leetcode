package leetcode

import "sort"

// Interval is an inclusive interval
type Interval struct {
	Start int
	End   int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start || intervals[i].Start == intervals[j].Start && intervals[i].End < intervals[j].End
	})
	r := []Interval{}
	
	for _, ival := range intervals {
		last := len(r) - 1
		if last < 0 || r[last].End < ival.Start {
			r = append(r, ival)
		} else {
			r[last] = Interval{minInt(r[last].Start, ival.Start), -minInt(-ival.End, -r[last].End)}
		}
	}
	return r
}
